package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func HandleHome(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}

func UserPostRequest(w http.ResponseWriter, request *http.Request) {
	user := new(User)
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %d", err)
		return
	}

	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %d", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
