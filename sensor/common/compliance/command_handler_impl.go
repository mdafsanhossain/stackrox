package compliance

import (
	"errors"

	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/internalapi/compliance"
	"github.com/stackrox/rox/generated/internalapi/sensor"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/sensor/common/roxmetadata"
)

var (
	log = logging.LoggerForModule()
)

type commandHandlerImpl struct {
	roxMetadata roxmetadata.Metadata

	commands chan *central.ScrapeCommand
	updates  chan *central.ScrapeUpdate

	service Service

	scrapeIDToState map[string]*scrapeState

	stopC    concurrency.ErrorSignal
	stoppedC concurrency.ErrorSignal
}

func (c *commandHandlerImpl) Start() {
	go c.run()
}

func (c *commandHandlerImpl) Stop(err error) {
	c.stopC.SignalWithError(err)
}

func (c *commandHandlerImpl) Stopped() concurrency.ReadOnlyErrorSignal {
	return &c.stoppedC
}

func (c *commandHandlerImpl) SendCommand(command *central.ScrapeCommand) bool {
	select {
	case c.commands <- command:
		return true
	case <-c.stoppedC.Done():
		return false
	}
}

func (c *commandHandlerImpl) Output() <-chan *central.ScrapeUpdate {
	return c.updates
}

func (c *commandHandlerImpl) run() {
	defer c.stoppedC.Signal()

	for {
		select {
		case <-c.stopC.Done():
			c.stoppedC.SignalWithError(c.stopC.Err())

		case command, ok := <-c.commands:
			if !ok {
				c.stoppedC.SignalWithError(errors.New("scrape command input closed"))
				return
			}
			if command.GetScrapeId() == "" {
				log.Errorf("received a command with no id: %s", proto.MarshalTextString(command))
				continue
			}
			if update := c.runCommand(command); update != nil {
				c.sendUpdate(update)
			}

		case result, ok := <-c.service.Output():
			if !ok {
				c.stoppedC.SignalWithError(errors.New("compliance return input closed"))
				return
			}
			if updates := c.commitResult(result); len(updates) > 0 {
				c.sendUpdates(updates)
			}
		}
	}
}

func (c *commandHandlerImpl) runCommand(command *central.ScrapeCommand) *central.ScrapeUpdate {
	switch command.Command.(type) {
	case *central.ScrapeCommand_StartScrape:
		return c.startScrape(command.GetScrapeId(), command.GetStartScrape().GetHostnames())
	case *central.ScrapeCommand_KillScrape:
		return c.killScrape(command.GetScrapeId())
	default:
		log.Errorf("unrecognized scrape command: %s", proto.MarshalTextString(command))
	}
	return nil
}

func (c *commandHandlerImpl) startScrape(scrapeID string, expectedHosts []string) *central.ScrapeUpdate {
	// Check that the scrape is not already running.
	if _, running := c.scrapeIDToState[scrapeID]; running {
		return nil
	}

	numResults := c.service.RunScrape(&sensor.MsgToCompliance{
		Msg: &sensor.MsgToCompliance_Trigger{
			Trigger: &sensor.MsgToCompliance_TriggerRun{
				ScrapeId: scrapeID,
			},
		},
	})

	// If we succeeded, start tracking the scrape and send a message to central.
	c.scrapeIDToState[scrapeID] = newScrapeState(scrapeID, numResults, expectedHosts)
	log.Infof("started scrape %q with %d results desired", scrapeID, numResults)
	return scrapeStarted(scrapeID, "")
}

func (c *commandHandlerImpl) killScrape(scrapeID string) *central.ScrapeUpdate {
	//// If killed successfully, remove the scrape from tracking.
	delete(c.scrapeIDToState, scrapeID)
	return scrapeKilled(scrapeID, "")
}

func (c *commandHandlerImpl) commitResult(result *compliance.ComplianceReturn) (ret []*central.ScrapeUpdate) {
	// Check that the scrape has not already been killed.
	scrapeState, running := c.scrapeIDToState[result.GetScrapeId()]
	if !running {
		log.Errorf("received result for scrape not tracked: %q", result.GetScrapeId())
		return
	}

	// Check that we have not already received a result for the host.
	if scrapeState.foundNodes.Contains(result.GetNodeName()) {
		log.Errorf("received duplicate result in scrape %s for node %s", result.GetScrapeId(), result.GetNodeName())
		return
	}
	scrapeState.desiredNodes--

	// Check if the node did not exist in the scrape request
	if !scrapeState.remainingNodes.Contains(result.GetNodeName()) {
		log.Errorf("found node %s not requested by Central for scrape %s", result.GetNodeName(), result.GetScrapeId())
		return
	}

	// Pass the update back to central.
	scrapeState.remainingNodes.Remove(result.GetNodeName())
	ret = append(ret, scrapeUpdate(result))
	// If that was the last expected update, kill the scrape.
	if scrapeState.desiredNodes == 0 || scrapeState.remainingNodes.Cardinality() == 0 {
		if scrapeState.remainingNodes.Cardinality() != 0 {
			log.Warnf("compliance data for the following nodes was not collected: %+v", scrapeState.remainingNodes.AsSlice())
		}
		if update := c.killScrape(result.GetScrapeId()); update != nil {
			ret = append(ret, update)
		}
	}

	return
}

func (c *commandHandlerImpl) sendUpdates(updates []*central.ScrapeUpdate) {
	if len(updates) > 0 {
		for _, update := range updates {
			c.sendUpdate(update)
		}
	}
}

func (c *commandHandlerImpl) sendUpdate(update *central.ScrapeUpdate) {
	select {
	case <-c.stoppedC.Done():
		log.Errorf("failed to send update: %s", proto.MarshalTextString(update))
		return
	case c.updates <- update:
		return
	}
}

// Helper functions.
///////////////////

func scrapeStarted(scrapeID, err string) *central.ScrapeUpdate {
	return &central.ScrapeUpdate{
		ScrapeId: scrapeID,
		Update: &central.ScrapeUpdate_ScrapeStarted{
			ScrapeStarted: &central.ScrapeStarted{
				ErrorMessage: err,
			},
		},
	}
}

func scrapeKilled(scrapeID, err string) *central.ScrapeUpdate {
	return &central.ScrapeUpdate{
		ScrapeId: scrapeID,
		Update: &central.ScrapeUpdate_ScrapeKilled{
			ScrapeKilled: &central.ScrapeKilled{
				ErrorMessage: err,
			},
		},
	}
}

func scrapeUpdate(result *compliance.ComplianceReturn) *central.ScrapeUpdate {
	return &central.ScrapeUpdate{
		ScrapeId: result.GetScrapeId(),
		Update: &central.ScrapeUpdate_ComplianceReturn{
			ComplianceReturn: result,
		},
	}
}
