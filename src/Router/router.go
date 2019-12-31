package Router

import (
	"Routers"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return Routers.SetupRoutes(r)
}
