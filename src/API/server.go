package API

import (
	"Router"
	"net/http"
)

func Run() {
	router := Router.New()
	http.ListenAndServe(":9090", router)
}
