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
