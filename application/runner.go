package application

import (
	"os"
	"os/signal"

	"github.com/redhaanggara21/docker-go-rabbitmq/util"
)

func AppRunner(daemon util.Daemon) error {
	err := daemon.Start()
	if err != nil {
		return err
	}

	osSignals := make(chan os.Signal)
	signal.Notify(osSignals, os.Interrupt)

	select {
	case <-osSignals:
		util.Log.Infof("osSignal Interrupt trigerred")
		return daemon.Stop()
	}
}
