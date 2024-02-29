package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func readJSON(req *http.Request, target interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(&target); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}
