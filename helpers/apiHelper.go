package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/golobby/container"
	"test.com/apiTest/configurations"
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

//RootHandler handles requests and errors
type RootHandler func(http.ResponseWriter, *http.Request) error

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error occured: %v", err) // Log the error.

	JsonBadRequest(w, err.Error)
}

//AuthenticationMiddleware checks if token is valid
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		bearer := request.Header.Get("Authorization")
		jwtToken := strings.Split(bearer, " ")
		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(jwtToken[1], claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			config := configurations.Configuration{}
			container.Make(&config)
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(config.Secret), nil
		})

		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("Unauthorized")
			fmt.Errorf(err.Error())
			return
		}

		if token.Valid {
			request.Header.Add("userId", claims.Subject)
			next.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("Unauthorized")
		}

	})
}
