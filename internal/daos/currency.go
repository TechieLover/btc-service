package daos

type ICurrencyDaos interface {
	GetCurrency(k string) (interface{}, bool)
}

type CurrencyInfoDaos struct {
}

func NewCurrency() ICurrencyDaos {
	return &CurrencyInfoDaos{}
}

func (c *CurrencyInfoDaos) GetCurrency(k string) (interface{}, bool) {
	ch := New()
	return ch.Get(k)
}
