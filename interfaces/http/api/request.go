package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// BindBody bind struct of request body.
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

// BindQuery bind struct of query strings.
func BindQuery(r *http.Request, b interface{}) error {
	m := make(map[string]interface{})
	v := r.URL.Query()
	if v == nil {
		return nil
	}
	for key, vs := range v {
		m[key] = vs[0]
	}

	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(j, b); err != nil {
		return err
	}

	return nil
}
