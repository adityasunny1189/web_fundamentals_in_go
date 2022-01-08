package handler

import (
	"net/http"

	"github.com/adityasunny1189/web_fundamentals_in_go/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
