// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// errInvalidToken is returned when the prometheus token is invalid.
var errInvalidToken = errors.New("Invalid or missing prometheus token")

// Handler is an http Metrics handler.
type Handler struct {
	metrics http.Handler
	token   string
}

// New returns a new metrics handler.
func New(token string) *Handler {
	return &Handler{
		metrics: promhttp.Handler(),
		token:   token,
	}
}

// ServeHTTP responds to an http.Request and writes system
// metrics to the response body in plain text format.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// extracts the bearer token from the request.
	token := r.Header.Get("Authorization")
	token = strings.TrimSpace(token)
	token = strings.TrimPrefix(token, "Bearer ")

	switch {
	case token != h.token:
		http.Error(w, errInvalidToken.Error(), 401)
	default:
		h.metrics.ServeHTTP(w, r)
	}
}
