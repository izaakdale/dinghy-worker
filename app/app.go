package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/izaakdale/dinghy-worker/consensus"
	"github.com/izaakdale/dinghy-worker/discovery"
	"github.com/kelseyhightower/envconfig"
)

var spec Specification

type Specification struct {
	discoveryCfg discovery.Config
	consensusCfg consensus.Config
}

type App struct{}

func New() *App {
	if err := envconfig.Process("", &spec.discoveryCfg); err != nil {
		log.Fatalf("failed to process discovery env vars: %v", err)
	}
	if err := envconfig.Process("", &spec.consensusCfg); err != nil {
		log.Fatalf("failed to process consensus env vars: %v", err)
	}
	return &App{}
}

func (a *App) Run() {
	consensus.New(spec.consensusCfg)

	node, evCh, err := discovery.NewMembership(spec.discoveryCfg)
	defer node.Leave()
	if err != nil {
		log.Fatal(err)
	}

	shCh := make(chan os.Signal, 2)
	signal.Notify(shCh, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-shCh:
			err := node.Leave()
			if err != nil {
				log.Fatalf("error leaving cluster %v", err)
			}
			os.Exit(1)
		case <-evCh:
			log.Println("event channel triggered")
		}
	}
}
