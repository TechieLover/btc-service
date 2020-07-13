package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

// GetRouter creates a router and registers all the routes for the
// service and returns it.
func GetRouter() http.Handler {
	router := httprouter.New()
	router.PanicHandler = PanicHandler
	setPingRoutes(router)
	setCurrencyRoutes(router)
	return router
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	log.Printf("Recovering from panic, Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
