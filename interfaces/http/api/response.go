package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
)

// WithFunc is set option parameter function.
type WithFunc func(h H)

// With set key value parameter.
func With(k string, v interface{}) WithFunc {
	return func(h H) {
		h[k] = v
	}
}

// WithMessage set message.
func WithMessage(m string) WithFunc {
	return With("message", m)
}

// WithError set error.
func WithError(err error) []WithFunc {
	return []WithFunc{
		WithMessage(err.Error()),
		func(h H) {
			switch e := err.(type) {
			case govalidator.Errors:
				errs := make(map[string]interface{})
				for _, e := range e.Errors() {
					switch e := e.(type) {
					case govalidator.Error:
						errs[e.Name] = H{
							"message": e.Error(),
							"type":    e.Validator,
						}
					default:
						return
					}
				}
				h["errors"] = errs
			default:
				return
			}
		},
	}
}

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

// BadRequest write bad request json body.
func BadRequest(w http.ResponseWriter, opts ...WithFunc) {
	m := H{"message": http.StatusText(http.StatusBadRequest)}
	for _, opt := range opts {
		opt(m)
	}

	JSON(w, &m, http.StatusBadRequest)
}
