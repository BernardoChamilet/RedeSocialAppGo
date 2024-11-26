package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(proxima http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		proxima(w, r)
	}
}

// Autenticar verifica a existência de cookies
func Autenticar(proxima http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//se n tiver cookies redireciona para /login
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		//se não segue para proxima func q devera ser a func da rota requisitada
		proxima(w, r)
	}
}
