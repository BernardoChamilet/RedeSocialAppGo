package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar gera um roteador com todas rotas configuradas
func Gerar() *mux.Router {
	return rotas.Configurar(mux.NewRouter())
}
