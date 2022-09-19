package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexdyukov/go-example-adserver/internal/creativehandler"
)

func main() {
	conf := GetCreativeConfig()

	handler := creativehandler.New(conf.ResponseWindow.Int64(), conf.PriceWindow.Int64())

	server := &http.Server{
		Addr:              conf.ServerAddress.String(),
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       time.Duration(conf.ResponseWindow) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(conf.ResponseWindow) * time.Millisecond,
		WriteTimeout:      time.Duration(2*conf.ResponseWindow) * time.Millisecond,
		IdleTimeout:       time.Duration(10*conf.ResponseWindow) * time.Millisecond,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ErrorLog:          nil,
		ConnState:         nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	log.Fatal(server.ListenAndServe())
}
