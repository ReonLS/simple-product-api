package middleware

import (
	"context"
	"net/http"
	"simple-product-api/utils"
	"strings"
)

//Happens first
func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization") //Bearer <Tokenstring>
		if authHeader == "" {
			http.Error(w, "Bearer String Empty", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			http.Error(w, "Not JWT", http.StatusUnauthorized)
			return
		}
		token := parts[1] //<jwt Token>
		claims, err := utils.ParseToken(token)
		if err != nil {
			http.Error(w, "Failed Parsing", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), utils.ClaimsKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//happens after
func AuthenticateRole(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		//Alur ambil claims, cek role
		claims, ok := utils.GetClaimsFromContext(r.Context())
		if !ok {
			http.Error(w, "Claims Failed", http.StatusUnauthorized)
			return
		}

		//cek roles
		if claims.Role != string(utils.RoleAdmin){
			http.Error(w, "Wrong Access Level", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}