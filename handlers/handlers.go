package handlers

import (
	"net/http"

	"github.com/ItsAchance/what2eat/components"
)

func Startpage(w http.ResponseWriter, r *http.Request) {
	component := components.Base("You nerd")
	component.Render(r.Context(), w)
}
