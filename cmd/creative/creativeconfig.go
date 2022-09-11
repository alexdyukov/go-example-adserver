package main

import (
	"flag"
	"log"
	"sync"

	env "github.com/caarlos0/env/v6"
)

type CreativeConfig struct {
	ServerAddress  ServerAddress `env:"SERVER_ADDRESS" envDefault:":8080" envExpand:"true"`
	once           sync.Once
	ResponseWindow int64 `env:"RESPONSE_WINDOW" envDefault:"1000" envExpand:"true"`
	IDToRedirect   int64 `env:"ID_TO_REDIRECT" envDefault:"0" envExpand:"true"`
	PriceWindow    int64 `env:"PRICE_WINDOW" envDefault:"1000" envExpand:"true"`
}

//nolint
var creativeConfig CreativeConfig

func GetCreativeConfig() *CreativeConfig {
	creativeConfig.once.Do(func() {
		if err := env.Parse(&creativeConfig); err != nil {
			log.Fatal(err)
		}

		flag.Int64Var(&creativeConfig.IDToRedirect, "i", creativeConfig.IDToRedirect, "id of client site")
		flag.Int64Var(&creativeConfig.PriceWindow, "p", creativeConfig.PriceWindow, "random price of request in [1; this) windows")
		flag.Int64Var(&creativeConfig.ResponseWindow, "r", creativeConfig.ResponseWindow, "random response in [0; this) window")
		flag.Var(&creativeConfig.ServerAddress, "a", "http listen address ")

		flag.Parse()
	})

	return &creativeConfig
}
