package main

import (
	logger "github.com/sirupsen/logrus"
)

// @title Vehicle Management Service
// @version 1.0
// @description Vehicle Management Service.
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url https://github.com/kalilventura/vehicle-management-service
// @contact.email kalilventur@gmail.com
//
// @license.name MIT License
// @license.url https://opensource.org/license/mit
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
