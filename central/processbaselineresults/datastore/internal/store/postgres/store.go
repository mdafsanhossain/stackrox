// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/sync"
	"gorm.io/gorm"
)

const (
	baseTable = "process_baseline_results"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000

	cursorBatchSize = 50
	deleteBatchSize = 5000
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.ProcessBaselineResultsSchema
	targetResource = resources.DeploymentExtension
)

// Store is the interface to interact with the storage for storage.ProcessBaselineResults
type Store interface {
	Upsert(ctx context.Context, obj *storage.ProcessBaselineResults) error
	UpsertMany(ctx context.Context, objs []*storage.ProcessBaselineResults) error
	Delete(ctx context.Context, deploymentId string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) error
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, deploymentId string) (bool, error)

	Get(ctx context.Context, deploymentId string) (*storage.ProcessBaselineResults, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storage.ProcessBaselineResults, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storage.ProcessBaselineResults, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storage.ProcessBaselineResults) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db    *postgres.DB
	mutex sync.RWMutex
}

// New returns a new Store instance using the provided sql instance.
func New(db *postgres.DB) Store {
	return &storeImpl{
		db: db,
	}
}

//// Helper functions

func insertIntoProcessBaselineResults(ctx context.Context, batch *pgx.Batch, obj *storage.ProcessBaselineResults) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(obj.GetDeploymentId()),
		pgutils.NilOrUUID(obj.GetClusterId()),
		obj.GetNamespace(),
		serialized,
	}

	finalStr := "INSERT INTO process_baseline_results (DeploymentId, ClusterId, Namespace, serialized) VALUES($1, $2, $3, $4) ON CONFLICT(DeploymentId) DO UPDATE SET DeploymentId = EXCLUDED.DeploymentId, ClusterId = EXCLUDED.ClusterId, Namespace = EXCLUDED.Namespace, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	return nil
}

func (s *storeImpl) copyFromProcessBaselineResults(ctx context.Context, tx *postgres.Tx, objs ...*storage.ProcessBaselineResults) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{

		"deploymentid",

		"clusterid",

		"namespace",

		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{

			pgutils.NilOrUUID(obj.GetDeploymentId()),

			pgutils.NilOrUUID(obj.GetClusterId()),

			obj.GetNamespace(),

			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetDeploymentId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = nil

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"process_baseline_results"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*postgres.Conn, func(), error) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Release, nil
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.ProcessBaselineResults) error {
	conn, release, err := s.acquireConn(ctx, ops.Get, "ProcessBaselineResults")
	if err != nil {
		return err
	}
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromProcessBaselineResults(ctx, tx, objs...); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.ProcessBaselineResults) error {
	conn, release, err := s.acquireConn(ctx, ops.Get, "ProcessBaselineResults")
	if err != nil {
		return err
	}
	defer release()

	for _, obj := range objs {
		batch := &pgx.Batch{}
		if err := insertIntoProcessBaselineResults(ctx, batch, obj); err != nil {
			return err
		}
		batchResults := conn.SendBatch(ctx, batch)
		var result *multierror.Error
		for i := 0; i < batch.Len(); i++ {
			_, err := batchResults.Exec()
			result = multierror.Append(result, err)
		}
		if err := batchResults.Close(); err != nil {
			return err
		}
		if err := result.ErrorOrNil(); err != nil {
			return err
		}
	}
	return nil
}

//// Helper functions - END

//// Interface functions

// Upsert saves the current state of an object in storage.
func (s *storeImpl) Upsert(ctx context.Context, obj *storage.ProcessBaselineResults) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "ProcessBaselineResults")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource).
		ClusterID(obj.GetClusterId()).Namespace(obj.GetNamespace())
	if !scopeChecker.IsAllowed() {
		return sac.ErrResourceAccessDenied
	}

	return pgutils.Retry(func() error {
		return s.upsert(ctx, obj)
	})
}

// UpsertMany saves the state of multiple objects in the storage.
func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.ProcessBaselineResults) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "ProcessBaselineResults")

	scopeChecker := sac.GlobalAccessScopeChecker(ctx).AccessMode(storage.Access_READ_WRITE_ACCESS).Resource(targetResource)
	if !scopeChecker.IsAllowed() {
		var deniedIDs []string
		for _, obj := range objs {
			subScopeChecker := scopeChecker.ClusterID(obj.GetClusterId()).Namespace(obj.GetNamespace())
			if !subScopeChecker.IsAllowed() {
				deniedIDs = append(deniedIDs, obj.GetDeploymentId())
			}
		}
		if len(deniedIDs) != 0 {
			return errors.Wrapf(sac.ErrResourceAccessDenied, "modifying processBaselineResultss with IDs [%s] was denied", strings.Join(deniedIDs, ", "))
		}
	}

	return pgutils.Retry(func() error {
		// Lock since copyFrom requires a delete first before being executed.  If multiple processes are updating
		// same subset of rows, both deletes could occur before the copyFrom resulting in unique constraint
		// violations
		if len(objs) < batchAfter {
			s.mutex.RLock()
			defer s.mutex.RUnlock()

			return s.upsert(ctx, objs...)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()

		return s.copyFrom(ctx, objs...)
	})
}

// Delete removes the object associated to the specified ID from the store.
func (s *storeImpl) Delete(ctx context.Context, deploymentId string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "ProcessBaselineResults")

	q := search.ConjunctionQuery(
		search.NewQueryBuilder().AddDocIDs(deploymentId).ProtoQuery(),
	)

	return pgSearch.RunDeleteRequestForSchema(ctx, schema, q, s.db)
}

// DeleteByQuery removes the objects from the store based on the passed query.
func (s *storeImpl) DeleteByQuery(ctx context.Context, query *v1.Query) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "ProcessBaselineResults")

	return pgSearch.RunDeleteRequestForSchema(ctx, schema, query, s.db)
}

// DeleteMany removes the objects associated to the specified IDs from the store.
func (s *storeImpl) DeleteMany(ctx context.Context, identifiers []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "ProcessBaselineResults")

	// Batch the deletes
	localBatchSize := deleteBatchSize
	numRecordsToDelete := len(identifiers)
	for {
		if len(identifiers) == 0 {
			break
		}

		if len(identifiers) < localBatchSize {
			localBatchSize = len(identifiers)
		}

		identifierBatch := identifiers[:localBatchSize]
		q := search.NewQueryBuilder().AddDocIDs(identifierBatch...).ProtoQuery()

		if err := pgSearch.RunDeleteRequestForSchema(ctx, schema, q, s.db); err != nil {
			return errors.Wrapf(err, "unable to delete the records.  Successfully deleted %d out of %d", numRecordsToDelete-len(identifiers), numRecordsToDelete)
		}

		// Move the slice forward to start the next batch
		identifiers = identifiers[localBatchSize:]
	}

	return nil
}

// Count returns the number of objects in the store.
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "ProcessBaselineResults")

	return pgSearch.RunCountRequestForSchema(ctx, schema, search.EmptyQuery(), s.db)
}

// Exists returns if the ID exists in the store.
func (s *storeImpl) Exists(ctx context.Context, deploymentId string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "ProcessBaselineResults")

	q := search.ConjunctionQuery(
		search.NewQueryBuilder().AddDocIDs(deploymentId).ProtoQuery(),
	)

	count, err := pgSearch.RunCountRequestForSchema(ctx, schema, q, s.db)
	// With joins and multiple paths to the scoping resources, it can happen that the Count query for an object identifier
	// returns more than 1, despite the fact that the identifier is unique in the table.
	return count > 0, err
}

// Get returns the object, if it exists from the store.
func (s *storeImpl) Get(ctx context.Context, deploymentId string) (*storage.ProcessBaselineResults, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "ProcessBaselineResults")

	q := search.ConjunctionQuery(
		search.NewQueryBuilder().AddDocIDs(deploymentId).ProtoQuery(),
	)

	data, err := pgSearch.RunGetQueryForSchema[storage.ProcessBaselineResults](ctx, schema, q, s.db)
	if err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	return data, true, nil
}

// GetByQuery returns the objects from the store matching the query.
func (s *storeImpl) GetByQuery(ctx context.Context, query *v1.Query) ([]*storage.ProcessBaselineResults, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetByQuery, "ProcessBaselineResults")

	rows, err := pgSearch.RunGetManyQueryForSchema[storage.ProcessBaselineResults](ctx, schema, query, s.db)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return rows, nil
}

// GetMany returns the objects specified by the IDs from the store as well as the index in the missing indices slice.
func (s *storeImpl) GetMany(ctx context.Context, identifiers []string) ([]*storage.ProcessBaselineResults, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "ProcessBaselineResults")

	if len(identifiers) == 0 {
		return nil, nil, nil
	}

	q := search.ConjunctionQuery(
		search.NewQueryBuilder().AddDocIDs(identifiers...).ProtoQuery(),
	)

	rows, err := pgSearch.RunGetManyQueryForSchema[storage.ProcessBaselineResults](ctx, schema, q, s.db)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			missingIndices := make([]int, 0, len(identifiers))
			for i := range identifiers {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	resultsByID := make(map[string]*storage.ProcessBaselineResults, len(rows))
	for _, msg := range rows {
		resultsByID[msg.GetDeploymentId()] = msg
	}
	missingIndices := make([]int, 0, len(identifiers)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input identifiers
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.ProcessBaselineResults, 0, len(resultsByID))
	for i, identifier := range identifiers {
		if result, ok := resultsByID[identifier]; !ok {
			missingIndices = append(missingIndices, i)
		} else {
			elems = append(elems, result)
		}
	}
	return elems, missingIndices, nil
}

// GetIDs returns all the IDs for the store.
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.ProcessBaselineResultsIDs")

	result, err := pgSearch.RunSearchRequestForSchema(ctx, schema, search.EmptyQuery(), s.db)
	if err != nil {
		return nil, err
	}

	identifiers := make([]string, 0, len(result))
	for _, entry := range result {
		identifiers = append(identifiers, entry.ID)
	}

	return identifiers, nil
}

// Walk iterates over all of the objects in the store and applies the closure.
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.ProcessBaselineResults) error) error {
	fetcher, closer, err := pgSearch.RunCursorQueryForSchema[storage.ProcessBaselineResults](ctx, schema, search.EmptyQuery(), s.db)
	if err != nil {
		return err
	}
	defer closer()
	for {
		rows, err := fetcher(cursorBatchSize)
		if err != nil {
			return pgutils.ErrNilIfNoRows(err)
		}
		for _, data := range rows {
			if err := fn(data); err != nil {
				return err
			}
		}
		if len(rows) != cursorBatchSize {
			break
		}
	}
	return nil
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed.
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed.
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}

//// Interface functions - END

//// Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db *postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db *postgres.DB) {
	dropTableProcessBaselineResults(ctx, db)
}

func dropTableProcessBaselineResults(ctx context.Context, db *postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS process_baseline_results CASCADE")

}

//// Used for testing - END
