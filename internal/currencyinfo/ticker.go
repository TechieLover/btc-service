package currencyinfo

import (
	"vdart/btc-service/dtos"
	"vdart/btc-service/internal/adapters"
	"vdart/btc-service/internal/daos"

	cache "github.com/patrickmn/go-cache"
)

func SetCurrencyInfoAndSymbol(avaicurr []string, client *adapters.WSClient, ch *daos.Pool) {
	curResp := dtos.CurrencyResponse{}
	for i := 0; i < len(avaicurr); i++ {
		symResp, _ := client.GetSymbol(avaicurr[i])
		currinfo, _ := client.GetCurrencyInfo(symResp.BaseCurrency)
		curResp.FeeCurrency = symResp.FeeCurrency
		curResp.ID = currinfo.ID
		curResp.FullName = currinfo.FullName
		ch.Set(avaicurr[i], curResp, cache.NoExpiration)
	}
}
