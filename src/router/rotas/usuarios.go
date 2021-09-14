package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:    "/maxsum",
		Metodo: http.MethodPost,
		Funcao: controllers.MaxSum,
	},
}
