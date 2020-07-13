package services

import (
	"vdart/btc-service/internal/daos"
)

type ICurrencyInfo interface {
	GetCurrency(k string) (interface{}, bool)
}

type CurrencyInfo struct {
	cur daos.ICurrencyDaos
}

func New() ICurrencyInfo {
	return &CurrencyInfo{
		cur: daos.NewCurrency(),
	}
}

func (c *CurrencyInfo) GetCurrency(k string) (interface{}, bool) {
	return c.cur.GetCurrency(k)
}
