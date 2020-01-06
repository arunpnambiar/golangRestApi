package apiController

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("Its Working.......!!!"))
}
func SaveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("Save Working.......!!!"))
}
