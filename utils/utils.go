package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func WriteINFO(w http.ResponseWriter, status int, key string, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(map[string]string{key: msg})
}

func ParseJSON(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

func ReadRequestID(r *http.Request) (int, error) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("missing account id")
	}

	accountID, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("internal server error")
	}
	return accountID, nil
}
