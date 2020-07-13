package handlers

import (
	"net/http"
	"strings"
	"vdart/btc-service/internal/config"
	"vdart/btc-service/internal/services"
	"vdart/btc-service/utilities"

	"github.com/julienschmidt/httprouter"
)

func setCurrencyRoutes(router *httprouter.Router) {
	router.GET("/v1/currency/:symbol", currencyInfo)
	router.GET("/v1/allcurrency", allcurrency)
}

func currencyInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	avaicurr := strings.Split(config.CURRENCY, ",")
	symb := strings.ToUpper(ps.ByName("symbol"))
	val := utilities.Contains(avaicurr, symb)
	if !val {
		writeJSONMessage("Please Provide a proper symbol", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	ser := services.New()
	foo, found := ser.GetCurrency(symb)
	if found {
		writeJSONStruct(foo, http.StatusOK, w)
		return
	}
	writeJSONMessage("Data not found", ERR_MSG, http.StatusBadRequest, w)
	return
}

func allcurrency(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ser := services.New()
	avaicurr := strings.Split(config.CURRENCY, ",")
	result := make([]interface{}, len(avaicurr))
	count := 0
	for i := 0; i < len(avaicurr); i++ {
		foo, found := ser.GetCurrency(avaicurr[i])
		if found {
			result[i] = foo
			count++
		}
	}
	if count == 0 {
		writeJSONMessage("Data not found", ERR_MSG, http.StatusBadRequest, w)
	}
	writeJSONStruct(result, http.StatusOK, w)
	return
}
