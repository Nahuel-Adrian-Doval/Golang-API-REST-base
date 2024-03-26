package router

import (
	"net/http"

	"github.com/Nahuel-Adrian-Doval/Golang-API-REST-base/internal/api/controllers"
)

func Routes(prefix string, mux *http.ServeMux) {

	mux.HandleFunc("GET "+prefix+"hello", controllers.Hello)
}
