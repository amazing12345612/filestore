package handler

import (
	"net/http"
)

func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.FormValue("username")
		token := r.FormValue("token")
		if len(username) < 3 || !IsTokenValid(token) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}
