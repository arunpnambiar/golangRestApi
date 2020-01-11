package Response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func response(w http.ResponseWriter, statuscode int, data interface{}) {
	w.WriteHeader(statuscode)
	err := json.NewEncoder(w).Encode(data)
	if err != nill {
		fmt.Fprint(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statuscode int, err error) {
	if err != nill {
		response(w, statuscode, struct {
			Error string `json:string`
		}{
			Error: err.Error(),
		})
		return
	}
	response(w, http.StatusBadRequest, nil)
}
