package main

import (
	"log"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/config"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/di"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	server.Start()

}
