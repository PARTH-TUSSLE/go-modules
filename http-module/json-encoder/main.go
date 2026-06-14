package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // res.type("json")
	w.WriteHeader(http.StatusOK) // res.status(200)

	res := map[string]any{
		"ok": true,
		"message": "JSON encoding successfull !",
		"datetime": time.Now().UTC(), 
	}

	_ = json.NewEncoder(w).Encode(res) // res.send(JSON.stringify(response))
}

func main() {

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}