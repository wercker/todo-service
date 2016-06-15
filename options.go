package main

import (
	"github.com/mies/todo-service/storage"
	"gopkg.in/redis.v3"
	"gopkg.in/urfave/cli.v1"
)

var (
	serverFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "service-port",
			Value: "5000",
			Usage: "TCP service port",
		},
		cli.StringFlag{
			Name:  "state-store",
			Value: "redis",
			Usage: "where to store state",
		},
	}
)

type GlobalOptions struct {
	ServicePort string
	StateStore  string
}

func flagsFrom(flagSets ...[]cli.Flag) []cli.Flag {
	all := []cli.Flag{}
	for _, flagSet := range flagSets {
		all = append(all, flagSet...)
	}
	return all
}

func GlobalFlags() []cli.Flag {
	return flagsFrom(serverFlags)
}

func ParseOpts(c *cli.Context) GlobalOptions {
	opts := GlobalOptions{}
	opts.ServicePort = c.GlobalString("service-port")
	opts.StateStore = c.GlobalString("state-store")
	return opts
}

func getStateStore(c *cli.Context, options GlobalOptions) (storage.StateStore, error) {
	switch options.StateStore {
	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     "redis-master:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		return storage.NewRedisStateStore(client), nil

	default:
		return storage.NewMemStore(), nil
	}
}
