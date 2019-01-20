package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BindBody bind struct of request body
func BindBody(r *http.Request, b interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, b); err != nil {
		return err
	}

	return nil
}
