package main

import (
	"log"

	config "../../pkg/config"
	proto "../../pkg/proto/auth"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

func initServiceConfig(c *cli.Context) {
	config.Source = c.String("config")
	config.FilePath = c.String("config_path")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
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

	proto.RegisterAuthHandler(service.Server(), new(Auth))

	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
