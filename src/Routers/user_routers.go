package Routers

import (
	"apiController"
	"net/http"
)

var userRouters = []Route{
	{url: "/", method: http.MethodGet, handler: apiController.GetUser},
}
