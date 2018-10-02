package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/sensor/common/networkflow/manager"
	"github.com/stackrox/rox/sensor/common/sensor"
	"github.com/stackrox/rox/sensor/kubernetes/enforcer"
	"github.com/stackrox/rox/sensor/kubernetes/listener"
	"github.com/stackrox/rox/sensor/kubernetes/orchestrator"
)

func main() {
	logger := logging.LoggerForModule()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	s := sensor.NewSensor(
		logger,
		listener.New(),
		enforcer.MustCreate(),
		orchestrator.MustCreate(),
		manager.Singleton(),
	)
	s.Start()

	for {
		select {
		case sig := <-sigs:
			logger.Infof("Caught %s signal", sig)
			s.Stop()
			return
		}
	}
}
