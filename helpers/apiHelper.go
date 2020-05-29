package helpers

import (
	"encoding/json"
	"net/http"
)

//JsonOK respond with json message to client with OK(200) status
func JsonOK(writer http.ResponseWriter, message interface{}) {
	messageJSON, error := json.Marshal(message)

	writer.Header().Set("Content-Type", "application/json")

	if error == nil {
		writer.WriteHeader(http.StatusOK)
		writer.Write(messageJSON)
	} else {
		JsonBadRequest(writer, error.Error())
	}
}

//JsonBadRequest respond with json message to client with BadRequest(400) status
func JsonBadRequest(writer http.ResponseWriter, message interface{}) {
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(message)
}
