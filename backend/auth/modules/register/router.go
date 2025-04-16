package register

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouter(router *mux.Router) {
	router.HandleFunc("/register", RegisterController).Methods(http.MethodPost)
}
