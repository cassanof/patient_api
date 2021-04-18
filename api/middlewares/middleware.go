package middlewares

import (
	"errors"
	"net/http"

	"github.com/elleven11/patient_api/api/auth"
	"github.com/elleven11/patient_api/api/responses"
)

func SetJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, req)
	}
}

func SetAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := auth.TokenValidate
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, req)
	}
}
