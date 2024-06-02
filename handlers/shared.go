package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

// Make takes in a handler and returns a handler function.
func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error(
				"HTTP handler error",
				"err", err,
				"path", r.URL.Path,
			)
		}
	}
}

// Render renders a templ component.
func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
