package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do something before the next handler
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request dose not contain an authorization header"})
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOCK"))

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "failed to get provider"})
			return
		}

		//verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
		_, err = verifier.Verify(r.Context(), tokenString)
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid tokenString"})
			return
		}

		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		emailInterface, ok := claims["email"]

		if !ok {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "email claim missing"})
			return
		}

		email, ok := emailInterface.(string)
		if !ok {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": fmt.Sprintf("email claim is not a string: %T", emailInterface)})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
		// Do something after the next handler
	})
}
