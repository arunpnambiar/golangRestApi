package Config

import (
	"log"
	"os"
	"strconv"
)

var (
	PORT = 0
)

func GetConfig() {
	var err error
	/*calling environment file to get value*/
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Printf(err.Error())
		PORT = 3000
	}
}
