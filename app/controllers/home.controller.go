package controllers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

var logger = logrus.New()

func Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	if err := render.HTML(w, http.StatusOK, "home", map[string]any{
		"title": "Home Page Title",
		"body":  "Home Page Description",
	}); err != nil {
		logger.Printf("failed load html file %s", err.Error())
	}
}
