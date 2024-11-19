package handlers

import (
	"e-book_webapp/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "templates/index.html", nil)
}
