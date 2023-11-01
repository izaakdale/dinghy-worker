package discovery

import (
	"io"
	"log"
	"strconv"

	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
)

type (
	Config struct {
		BindAddr      string `envconfig:"BIND_ADDR"`
		BindPort      string `envconfig:"BIND_PORT"`
		AdvertiseAddr string `envconfig:"ADVERTISE_ADDR"`
		AdvertisePort string `envconfig:"ADVERTISE_PORT"`
		ClusterAddr   string `envconfig:"CLUSTER_ADDR"`
		ClusterPort   string `envconfig:"CLUSTER_PORT"`
	}
	Tag struct {
		Key   string
		Value string
	}
)

func NewMembership(name string, cfg Config, tags ...Tag) (*serf.Serf, chan serf.Event, error) {
	// since there will be multiple workers, we need unique names
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = cfg.AdvertiseAddr

	var err error
	conf.MemberlistConfig.AdvertisePort, err = strconv.Atoi(cfg.AdvertisePort)
	if err != nil {
		return nil, nil, err
	}
	conf.MemberlistConfig.BindAddr = cfg.BindAddr
	conf.MemberlistConfig.BindPort, err = strconv.Atoi(cfg.BindPort)
	if err != nil {
		return nil, nil, err
	}
	conf.MemberlistConfig.ProtocolVersion = 3 // Version 3 enable the ability to bind different port for each agent
	conf.NodeName = name

	// prevent annoying serf and memberlist logs
	conf.MemberlistConfig.Logger = log.New(io.Discard, "", log.Flags())
	conf.Logger = log.New(io.Discard, "", log.Flags())

	t := make(map[string]string, len(tags))
	t["name"] = name
	for _, tag := range tags {
		t[tag.Key] = tag.Value
	}
	conf.Tags = t

	evCh := make(chan serf.Event)
	conf.EventCh = evCh

	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Couldn't create cluster")
	}

	_, err = cluster.Join([]string{cfg.ClusterAddr + ":" + cfg.ClusterPort}, true)
	if err != nil {
		log.Printf("Couldn't join cluster, starting own: %v\n", err)
	}

	return cluster, evCh, nil
}
