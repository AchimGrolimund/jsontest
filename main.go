package main

import (
	log "github.com/sirupsen/logrus"
	"jsontest/app"
)

func init() {
	// set default log config
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

var (
	// Version is provided by ldflags
	Version = "unspecified"
	// Build is provided by ldflags
	Build = "unspecified"
)

func main() {
	log.Infof("Version: %s, Build: %s", Version, Build)

	app.Start()

}
