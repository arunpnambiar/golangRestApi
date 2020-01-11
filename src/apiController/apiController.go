package apiController

import (
	"db"
	"fmt"
	"log"
	"models"
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
		log.Fatal(err)
	}
	err = connection.Debug().Model(&models.User{}).Create(&r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
}
