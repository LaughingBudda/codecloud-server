package auth

import (
	"fmt"
	"net/http"
	"os"
	
	"github.com/LaughingBudda/codecloud-server/constants"
	"github.com/auth0-community/auth0"
	jose "gopkg.in/square/go-jose.v2"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte(constants.JWT_SECRET)
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{constants.API_AUDIENCE}

		configuration := auth0.NewConfiguration(secretProvider, audience, constants.AUTH0_DOMAIN, jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
