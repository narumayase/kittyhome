package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "kittyhome/docs"
	pc "kittyhome/internal/pet/infrastructure/controller"
	ph "kittyhome/internal/pet/infrastructure/handler"
	pr "kittyhome/internal/pet/infrastructure/repository"
	tc "kittyhome/internal/toy/infrastructure/controller"
	th "kittyhome/internal/toy/infrastructure/handler"
	tr "kittyhome/internal/toy/infrastructure/repository"
	"net/http"
)

//go:generate swag init
//go:generate swag fmt

// main
// @title                     Kitty Home Boilerplate API
// @version                   0.0.1
// @description               This is a sample boilerplate REST API.
// @termsOfService            http://sarasa
// @contact.name              Naru Mayase
// @contact.url               http://www.sarasa
// @contact.email             sarasa@gmail.com
// @host                      localhost:8080
// @BasePath                  /
// @securityDefinitions.basic BasicAuthfunc
func main() {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	//r = r.PathPrefix("/api/v1").Subrouter()

	ph.AddRoutes(r, ph.Build(pc.Build(pr.Build())))
	th.AddRoutes(r, th.Build(tc.Build(tr.Build())))

	log.Err(http.ListenAndServe(":8080", r))
}
