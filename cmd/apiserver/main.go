package main

import (
	"flag"
	"log"

	"github.com/lehius/http-resy-api/internal/app/apiserver"
	"github.com/lehius/http-resy-api/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {
	flag.Parse()

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
