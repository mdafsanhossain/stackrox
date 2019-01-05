package compliancereturn

import (
	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/central/sensorevent/service/pipeline"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/pkg/logging"
)

var (
	log = logging.LoggerForModule()
)

// Template design pattern. We define control flow here and defer logic to subclasses.
//////////////////////////////////////////////////////////////////////////////////////

// NewPipeline returns a new instance of Pipeline.
func NewPipeline() pipeline.Pipeline {
	return &pipelineImpl{}
}

type pipelineImpl struct{}

// Run runs the pipeline template on the input and returns the output.
func (s *pipelineImpl) Run(event *central.SensorEvent, _ pipeline.EnforcementInjector) (err error) {
	// do nothing for now.
	log.Infof("ignoring compliance run: %s", proto.MarshalTextString(event))
	return nil
}
