package main

import (
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.Info("⚙️ Initializing application...")
	defer handlePanic()

	modules := injectModules()
	settings := injectSettings()

	StartServer(modules, settings)
}

func handlePanic() {
	if r := recover(); r != nil {
		logger.WithField("panic", r).
			Fatal("🚨 A critical and unrecoverable error occurred, forcing the application to stop.")
	}
}
