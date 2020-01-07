package apiController

import (
	"db"
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("Its Working.......!!!"))
}
func SaveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	connection, err := db.Connection()
	if err != nil {

	}

	w.Write([]byte("Save Working.......!!!"))
	defer connection.Close()
}
