package handlers

import (
	"net/http"

	"github.com/robwestbrook/go-mylife/views/auth"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
