package routes

import(
	"github.com/gorilla/mux"
	"gofiber/persona_controller"
)

func SetPersonaRoutes(router *mux.Router){
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", persona_controller.All).Methods("GET")
	subRoute.HandleFunc("/save", persona_controller.Insert).Methods("POST")
	subRoute.HandleFunc("/delete", persona_controller.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", persona_controller.Get).Methods("GET")
	subRoute.HandleFunc("/hola", persona_controller.Holamundo).Methods("GET")
}
