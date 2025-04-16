package login

import "github.com/gorilla/mux"

func main() {
	router := mux.NewRouter()
	controller := LoginController()
	router.HandleFunc("/login", controller)
}
