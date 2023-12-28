package Authentication

import (
	"awesomeProject/Utilities"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nlpodyssey/gotokenizers/models"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := request.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(writer, request)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHandler := request.Header.Get("Authentication")

		if tokenHandler == "" {
			response = Utilities.Message(false, "Missing auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			Utilities.Respond(writer, response)
			return
		}

		splitted := strings.Split(tokenHandler, " ")

		if len(splitted) != 2 {
			response = Utilities.Message(false, "Wrong authentication token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			Utilities.Respond(writer, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = Utilities.Message(false, "Invalid token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			Utilities.Respond(writer, response)
			return
		}

		if !token.Valid {
			response := Utilities.Message(false, "Token is not valid")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			Utilities.Respond(writer, response)
			return
		}

		fmt.Sprintf("User %", tk.Value)
		ctx := context.WithValue(request.Context(), "user", tk.ID)
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}
