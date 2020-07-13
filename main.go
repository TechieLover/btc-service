package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"vdart/btc-service/dtos"
	"vdart/btc-service/internal/adapters"
	"vdart/btc-service/internal/config"

	"vdart/btc-service/internal/currencyinfo"
	"vdart/btc-service/internal/daos"
	"vdart/btc-service/internal/handlers"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	client, err := adapters.NewWSClient()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	avaicurr := strings.Split(config.CURRENCY, ",")
	fmt.Println(avaicurr)

	ch := daos.New()
	currencyinfo.SetCurrencyInfoAndSymbol(avaicurr, client, ch)
	tickerFeed := make([]<-chan adapters.WSNotificationTickerResponse, len(avaicurr))

	for i := 0; i < len(avaicurr); i++ {
		tickerFeed[i], _ = client.SubscribeTicker(avaicurr[i])
	}

	interrupt := make(chan os.Signal, 1)
	for i := 0; i < len(avaicurr); i++ {
		go func(i int) {
			for {
				select {
				case ticker := <-tickerFeed[i]:
					val, _ := ch.Get(ticker.Symbol)
					val1, _ := json.Marshal(val)
					curResp := dtos.CurrencyResponse{}
					json.Unmarshal(val1, &curResp)
					curResp.Ask = ticker.Ask
					curResp.Bid = ticker.Bid
					curResp.High = ticker.High
					curResp.Last = ticker.Last
					curResp.Low = ticker.Low
					curResp.Open = ticker.Open
					ch.Set(ticker.Symbol, curResp, cache.NoExpiration)
				case <-interrupt:
					fmt.Println("No Data is availabe")
					return
				}
			}
		}(i)
	}
	http.ListenAndServe(":"+config.PORT, handlers.GetRouter())
}
