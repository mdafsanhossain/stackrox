package generatecerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stackrox/rox/pkg/apiparams"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/roxctl/common"
	"github.com/stackrox/rox/roxctl/common/download"
	"github.com/stackrox/rox/roxctl/common/environment"
	"github.com/stackrox/rox/roxctl/common/flags"
	util2 "github.com/stackrox/rox/roxctl/common/util"
	"github.com/stackrox/rox/roxctl/sensor/util"
)

func downloadCerts(env environment.Environment, outputDir, clusterIDOrName string, timeout time.Duration) error {
	clusterID, err := util.ResolveClusterID(env, clusterIDOrName, timeout)
	if err != nil {
		return err
	}

	body, err := json.Marshal(&apiparams.ClusterCertGen{ID: clusterID})
	if err != nil {
		return err
	}

	resp, err := util2.DoHTTPRequestAndCheck200(env, "/api/extensions/certgen/cluster", timeout, http.MethodPost, bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer utils.IgnoreError(resp.Body.Close)

	fileName, err := download.ParseFilenameFromHeader(resp.Header)
	if err != nil {
		fileName = fmt.Sprintf("cluster-%s-certs.yaml", clusterIDOrName)
		env.Logger().WarnfLn("could not obtain output file name from HTTP Response: %v. Defaulting to %q", err, fileName)
	}

	outputFileNameWithDir := filepath.Join(outputDir, fileName)
	createdFile, err := os.Create(outputFileNameWithDir)
	if err != nil {
		return errors.Wrap(err, "failed to create output file")
	}
	var fileClosed bool
	defer func() {
		if !fileClosed {
			_ = createdFile.Close()
		}
	}()

	_, err = io.Copy(createdFile, resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to write from response to file")
	}

	err = createdFile.Close()
	fileClosed = true
	if err != nil {
		return errors.Wrapf(err, "failed to close file at %s", outputFileNameWithDir)
	}
	env.Logger().InfofLn("Successfully downloaded new certs. Use kubectl apply -f %s to apply them.", outputFileNameWithDir)
	return nil
}

// Command defines the command.
func Command(cliEnvironment environment.Environment) *cobra.Command {
	var outputDir string

	c := &cobra.Command{
		Use:  "generate-certs <cluster-name-or-id>",
		Args: common.ExactArgsWithCustomErrMessage(1, "No cluster name or ID specified"),
		RunE: func(c *cobra.Command, args []string) error {
			if err := downloadCerts(cliEnvironment, outputDir, args[0], flags.Timeout(c)); err != nil {
				return errors.Wrap(err, "error downloading regenerated certs")
			}
			return nil
		},
	}

	c.PersistentFlags().StringVar(&outputDir, "output-dir", ".", "output directory for the YAML file")

	return c
}
