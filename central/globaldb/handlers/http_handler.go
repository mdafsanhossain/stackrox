package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	bolt "github.com/etcd-io/bbolt"
	"github.com/stackrox/rox/central/globaldb/export"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/osutils"
	"github.com/stackrox/rox/pkg/rocksdb"
	"github.com/stackrox/rox/pkg/utils"
)

var (
	log = logging.LoggerForModule()
)

const (
	dbFileFormat      = "stackrox_db_2006_01_02_15_04_05.zip"
	restoreFileFormat = "dbrestore_2006_01_02_15_04_05"
)

// BackupDB is a handler that writes a consistent view of the databases to the HTTP response.
func BackupDB(boltDB *bolt.DB, rocksDB *rocksdb.RocksDB) http.Handler {
	return serializeDB(boltDB, rocksDB)
}

func logAndWriteErrorMsg(w http.ResponseWriter, code int, t string, args ...interface{}) {
	errMsg := fmt.Sprintf(t, args...)
	log.Error(errMsg)
	http.Error(w, errMsg, code)
}

// This will EOF if we call exit at the end of a handler
func deferredRestart(ctx context.Context) {
	go func() {
		concurrency.WaitWithTimeout(ctx, 5*time.Second)
		time.Sleep(1 * time.Second)
		osutils.Restart()
	}()
}

// RestoreDB is a handler that takes in a DB and restores Central to it
func RestoreDB(boltDB *bolt.DB, rocksDB *rocksdb.RocksDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Info("Starting DB restore ...")
		filename := filepath.Join(os.TempDir(), time.Now().Format(restoreFileFormat))

		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			logAndWriteErrorMsg(w, http.StatusInternalServerError, "could not create temporary file for DB upload: %v", err)
			return
		}
		defer func() {
			_ = os.Remove(f.Name())
		}()
		defer utils.IgnoreError(f.Close)

		if _, err := io.Copy(f, req.Body); err != nil {
			_ = req.Body.Close()
			logAndWriteErrorMsg(w, http.StatusInternalServerError, "error storing upload in temporary location: %v", err)
			return
		}
		_ = req.Body.Close()
		if _, err := f.Seek(0, 0); err != nil {
			logAndWriteErrorMsg(w, http.StatusInternalServerError, "could not rewind to beginning of temporary file: %v", err)
			return
		}

		if err := export.Restore(f); err != nil {
			logAndWriteErrorMsg(w, http.StatusInternalServerError, "could not restore database backup: %v", err)
			return
		}
		log.Info("DB restore completed")

		// Now that we have verified the uploaded DB, close the current DB
		// and bounce Central

		if err := boltDB.Close(); err != nil {
			log.Errorf("unable to close bolt DB: %v", err)
		}
		rocksDB.Close()

		log.Info("Bouncing Central to pick up newly imported DB")
		deferredRestart(req.Context())
	})
}

func serializeDB(boltDB *bolt.DB, rocksDB *rocksdb.RocksDB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Info("Starting DB backup ...")
		filename := time.Now().Format(dbFileFormat)

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

		if err := export.Backup(req.Context(), boltDB, rocksDB, w); err != nil {
			logAndWriteErrorMsg(w, http.StatusInternalServerError, "could not create database backup: %v", err)
			return
		}
		log.Info("DB backup completed")
	}
}
