package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// BindBody bind struct of request body.
func BindBody(r *http.Request, b interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
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

// WrapAllowContentType wrap content type request.
func WrapAllowContentType(f http.HandlerFunc, types ...string) http.HandlerFunc {
	cT := []string{}
	for _, t := range types {
		cT = append(cT, strings.ToLower(t))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s := strings.ToLower(strings.TrimSpace(r.Header.Get("Content-Type")))
		if i := strings.Index(s, ";"); i > -1 {
			s = s[0:i]
		}

		for _, t := range cT {
			if t == s {
				f(w, r)
				return
			}
		}

		UnsupportedMediaType(w, WithMessage(fmt.Sprintf("Not supported type. Allow content-type: %s", cT)))
	}
}
