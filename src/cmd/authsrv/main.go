package main

import (
	"log"

	config "../../pkg/config"
	consulConfig "../../pkg/config/consul"
	proto "../../pkg/proto/auth"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:  "config",
				Usage: "Service configuration source, available options `consul` or `file`",
			},
			cli.StringFlag{
				Name:  "config_path",
				Usage: "Configuration file path, read only in case of `file` config source",
			},
		),
	)

	// TODO: config should be switchable based on environment variables or CLI options
	serviceConfig := consulConfig.InitConfig("auth/config/")
	if serviceConfig.Err != nil {
		log.Fatal(serviceConfig.Err)
		return
	}

	proto.RegisterAuthHandler(service.Server(), &Auth{config.InitConfig(serviceConfig)})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
