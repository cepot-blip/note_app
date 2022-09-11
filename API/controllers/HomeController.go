package controllers

import (
	"net/http"

	"github.com/cepot-blip/noteApp/API/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, `
	=========================================

	  W E L C O M E  T O  T H E  M Y  A P I

	=========================================
	`)
}
