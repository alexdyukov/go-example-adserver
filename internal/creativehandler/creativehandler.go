package creativehandler

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
	"time"
)

//easyjson:json
type response struct {
	ID    int64 `json:"id"`
	Price int64 `json:"price"`
}

func New(serverID, responseWindow, priceWindow int64) http.Handler {
	bigPriceWindow := big.NewInt(priceWindow)
	bigResponseWindow := big.NewInt(responseWindow)

	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		millisecondsToSleep, err := rand.Int(rand.Reader, bigResponseWindow)
		if err != nil {
			panic("crypto/rand internal error in sleep phase:" + err.Error())
		}

		time.Sleep(time.Duration(millisecondsToSleep.Int64()) * time.Millisecond)

		price, err := rand.Int(rand.Reader, bigPriceWindow)
		if err != nil {
			panic("crypto/rand internal error in price gen phase:" + err.Error())
		}

		res := response{
			ID:    serverID,
			Price: price.Int64(),
		}

		jsonData, err := json.Marshal(res)
		if err != nil {
			panic("cannot marshal response")
		}

		if _, err := responseWriter.Write(jsonData); err != nil {
			panic("cannot write response")
		}
	})
}
