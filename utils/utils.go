package utils

import (
	"encoding/json"
	"fly/repositories"
	"fmt"
	"net/http"
)

func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ValidarServicos() {
	repositories.CriarTabelas()
}