package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON write json body.
func JSON(w http.ResponseWriter, body interface{}, code int) error {
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(code)

	res, err := json.Marshal(body)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, string(res))
	return nil
}

// Notfound write notfound json body.
func Notfound(w http.ResponseWriter) {
	JSON(w, &H{"message": http.StatusText(http.StatusNotFound)}, http.StatusNotFound)
}

// InternalServerError write internal server error json body.
func InternalServerError(w http.ResponseWriter) {
	JSON(w, &H{"message": http.StatusText(http.StatusInternalServerError)}, http.StatusInternalServerError)
}
