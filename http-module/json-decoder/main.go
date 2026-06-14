package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON (w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_  = json.NewEncoder(w).Encode(data)
}

type TESTREQUEST struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}

func testHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok": false,
			"error": "Only POST is allowed !",
		})
		return
	}

	defer r.Body.Close()

	var REQ TESTREQUEST

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&REQ); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok": false,
			"error": "Invalid json format !",
		})
		return
	}

	REQ.Name = strings.TrimSpace(REQ.Name)

	if REQ.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok": false,
			"error": "Name cannot be empty !",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok": true,
		"data": REQ,
		"timestamp": time.Now().UTC(),
	})

}

func main () {

	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}