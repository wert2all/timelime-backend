package auth

import (
	"net/http"
	"strings"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// token := extractToken(req)
			// if token == "" {
			// 	http.Error(w, "Invalid token", http.StatusForbidden)
			// 	return
			// }

			// // Allow unauthenticated users in
			// if err != nil || c == nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			// userId, err := validateAndGetUserID(c)
			// if err != nil {
			// 	http.Error(w, "Invalid cookie", http.StatusForbidden)
			// 	return
			// }

			// // get the user from the database
			// user := getUserByID(db, userId)

			// // put it in context
			// ctx := context.WithValue(r.Context(), userCtxKey, user)

			// // and call the next with our new context
			// r = r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}

func extractToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	splitted := strings.Split(authHeader, " ")

	if len(splitted) == 2 {
		return splitted[1]
	} else {
		return ""
	}
}