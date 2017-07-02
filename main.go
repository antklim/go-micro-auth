package main

import (
	"log"

	config "./config"
	handler "./handler/auth"
	proto "./proto/auth"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

var (
	// Version of the service
	Version string

	// Build number of the service
	Build string
)

func initServiceConfig(c *cli.Context) {
	config.Source = c.String("config")
	config.FilePath = c.String("config_path")
}

func main() {
	log.Printf("Service Version: %s\n", Version)
	log.Printf("Service Build: %s\n", Build)

	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version(Version),
		micro.Flags(
			cli.StringFlag{
				Name:   "config",
				EnvVar: "CONFIG",
				Usage:  "Service configuration source, available options `consul` or `file`",
			},
			cli.StringFlag{
				Name:   "config_path",
				EnvVar: "CONFIG_PATH",
				Usage:  "Configuration file path, read only in case of `file` config source",
			},
		),
		micro.Action(initServiceConfig),
	)

	service.Init()

	// `config` Init() should be called after service init
	// All `config` variables are set on service init action
	configHandler, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterAuthHandler(service.Server(), &handler.Auth{ConfigHandler: configHandler})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
