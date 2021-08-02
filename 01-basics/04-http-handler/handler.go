package handler

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Name:       "james bond",
		Age:        32,
		Occupation: "secret agent",
	}
	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
