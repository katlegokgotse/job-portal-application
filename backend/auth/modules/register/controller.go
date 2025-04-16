package register

import (
	"encoding/json"
	"net/http"
)

func RegisterController(writer http.ResponseWriter, request *http.Request) {
	response := map[string]string{"messsage": "User has been registered"}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
