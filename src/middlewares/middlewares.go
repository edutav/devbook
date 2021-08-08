package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s - %s - %s", r.Method, r.Host, r.RequestURI)
		next(rw, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")
		next(rw, r)
	}
}
