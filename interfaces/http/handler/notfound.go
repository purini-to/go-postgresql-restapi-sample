package handler

import (
	"net/http"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/api"
)

var (
	notfoundStatus = http.StatusNotFound
)

// Notfound is notfound handler.
type Notfound struct{}

// Handler return notfound response.
func (n *Notfound) Handler(w http.ResponseWriter, r *http.Request) {
	api.JSON(w, &api.H{"message": http.StatusText(notfoundStatus)}, notfoundStatus)
}

// NewNotfound create notfound handler
func NewNotfound() *Notfound {
	return &Notfound{}
}
