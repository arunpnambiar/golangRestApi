package API

import (
	"Config"
	"Router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	Config.GetConfig()
	router := Router.New()
	PORT := Config.PORT
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
