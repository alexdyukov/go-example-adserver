package main

import (
	"flag"
	"log"
	"sync"

	"github.com/alexdyukov/go-example-adserver/internal/cliparams"
	env "github.com/caarlos0/env/v6"
)

type CreativeConfig struct {
	ServerAddress  cliparams.ServerAddress `env:"SERVER_ADDRESS" envDefault:":8080" envExpand:"true"`
	once           sync.Once
	ResponseWindow cliparams.UintWindow `env:"RESPONSE_WINDOW" envDefault:"1000" envExpand:"true"`
	PriceWindow    cliparams.UintWindow `env:"PRICE_WINDOW" envDefault:"1000" envExpand:"true"`
}

//nolint
var creativeConfig CreativeConfig

func GetCreativeConfig() *CreativeConfig {
	creativeConfig.once.Do(func() {
		if err := env.Parse(&creativeConfig); err != nil {
			log.Fatal(err)
		}

		flag.Var(&creativeConfig.PriceWindow, "p", "random price of request in [0; this) windows")
		flag.Var(&creativeConfig.ResponseWindow, "r", "random response in [0; this) window")
		flag.Var(&creativeConfig.ServerAddress, "a", "http listen address in 'addr:port' format")

		flag.Parse()
	})

	return &creativeConfig
}
