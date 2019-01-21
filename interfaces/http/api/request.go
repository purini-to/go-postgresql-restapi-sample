package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
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

// BindQuery bind struct of request query.
func BindQuery(r *http.Request, b interface{}) error {
	m := make(map[string]interface{})
	rctx := chi.RouteContext(r.Context())
	for k := len(rctx.URLParams.Keys) - 1; k >= 0; k-- {
		m[rctx.URLParams.Keys[k]] = rctx.URLParams.Values[k]
	}

	if len(m) == 0 {
		return nil
	}

	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, b)
	if err != nil {
		return err
	}
	return nil
}
