package utils

import (
	"os"
	"os/signal"
)

const (
	// GracefulShutdownCode - exit code
	GracefulShutdownCode = 2
)

// RegisterGracefulShutdown - sets handler to catch SIGTERM and SIGKILL signals
// after executing handle function, application exits with GracefulShutdownCode code
func RegisterGracefulShutdown(handle func(os.Signal)) {
	var signals = make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	go func() {
		var signal = <-signals
		handle(signal)
		os.Exit(GracefulShutdownCode)
	}()
}
