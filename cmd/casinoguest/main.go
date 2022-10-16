package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/CasinoTrade/CasinoGuest/internal/log"
	"github.com/CasinoTrade/CasinoGuest/internal/model/config"
	"github.com/CasinoTrade/CasinoGuest/internal/server"
	"github.com/CasinoTrade/CasinoGuest/internal/server/rest"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "none"
)

func main() {
	printVersion := flag.Bool("version", false, "Get version")
	flag.Parse()
	if *printVersion {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", date)
		fmt.Printf("Build Commit: %s\n", commit)
		return
	}

	// init
	logger := log.New(config.DefaultCfg().Logger)
	casino := server.New(logger.WithSource("base-server"))
	s := rest.New(config.DefaultCfg().Server, logger.WithSource("rest-server"), casino)

	// start
	casino.Start()
	s.Start()

	// shutdown handling
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, os.Interrupt)

	<-ch
	logger.Info("Handling interrupt")
	s.Stop()
	logger.Info("Done")
}
