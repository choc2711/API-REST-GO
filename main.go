package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gofiber/conexion"
	"gofiber/routes"
)


func main(){
	conexion.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr: ":80",
		Handler: router,
	}

	log.Println(server.ListenAndServe())
}