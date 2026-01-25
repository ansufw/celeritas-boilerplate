package handlers

import (
	"boilerplate/data"
	"net/http"

	"github.com/ansufw/celeritas"
)

type XMLPayload struct {
	ID      int64    `xml:"id"`
	Name    string   `xml:"name"`
	Hobbies []string `xml:"hobby"`
}

type Handlers struct {
	App    *celeritas.Celeritas
	Models *data.Model
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	// defer h.App.LoadTime(time.Now())
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering page: ", err)
	}
}
