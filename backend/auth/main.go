package main

import (
	"github.com/gorilla/mux"
	"main.go/modules/register"
)

func main() {
	router := *mux.NewRouter()
	register.RegisterRouter(&router)
	login.LoginRouter(&router)
}
