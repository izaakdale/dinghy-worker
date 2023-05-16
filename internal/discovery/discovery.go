package discovery

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
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
		Name          string `envconfig:"NAME"`
	}
	Tag struct {
		Key   string
		Value string
	}
)

func NewMembership(cfg Config, tags ...Tag) (*serf.Serf, chan serf.Event, error) {
	// since there will be multiple workers, we need unique names
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = cfg.AdvertiseAddr
	conf.MemberlistConfig.AdvertisePort, _ = strconv.Atoi(cfg.AdvertisePort)
	conf.MemberlistConfig.BindAddr = cfg.BindAddr
	conf.MemberlistConfig.BindPort, _ = strconv.Atoi(cfg.BindPort)
	conf.MemberlistConfig.ProtocolVersion = 3 // Version 3 enable the ability to bind different port for each agent

	name := fmt.Sprintf("%s-%s", cfg.Name, strings.Split(uuid.NewString(), "-")[0])
	conf.NodeName = name

	t := make(map[string]string, len(tags))
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
