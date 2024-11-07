package handler

import "net/http"

func Json(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}
