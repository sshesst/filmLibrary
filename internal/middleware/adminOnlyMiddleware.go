package middleware

import "net/http"

func AdminOnlyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value("UserRole").(bool)
		if !ok || isAdmin == false {
			http.Error(w, "В доступе отказано", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
