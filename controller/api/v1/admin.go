package controller

import (
	"net/http"
	"trip-planer/lib/render"
)

func Healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, nil, Health{"healthy"})
	}
}

type Health struct {
	Message string `json:"message"`
}
