package login

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", LoginController).Methods(http.MethodPost)
}
