package handlers

import (
	"errors"
	"net/http"
)

type validator struct {}

func (v *validator)ValidateJsonContentType(r *http.Request) error {
	content_type := r.Header.Get("Content-Type")

	if content_type != "application/json" {
		return errors.New("Content-Type must be application/json")
	}

	return nil
}