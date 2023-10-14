package persona_controller

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"gofiber/conexion"
	"gofiber/persona"
	"gofiber/handler"
	
)

func All(writer http.ResponseWriter,request *http.Request){
	personas := []persona.Persona{}
	db := conexion.Conexion()
	defer db.Close()

	db.Find(&personas)

	json, _ := json.Marshal(personas)
	handler.SendRespond(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter,request *http.Request){
 per := persona.Persona{}

 id:= mux.Vars(request)["id"]
 db := conexion.Conexion()
defer db.Close()

db.Find(&per,id)

if per.ID > 0 {
	json,_ := json.Marshal(per)
	handler.SendRespond(writer, http.StatusOK, json)
}else {
	handler.SendError(writer, http.StatusNotFound)
}

}

func Insert(writer http.ResponseWriter,request *http.Request){
	per := persona.Persona{}

	db := conexion.Conexion()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&per)

	if error != nil {
		log.Fatal(error)
		handler.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&per).Error

	if error != nil {
		log.Fatal(error)
		handler.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(per)

	handler.SendRespond(writer, http.StatusCreated,json)
}

func Delete(writer http.ResponseWriter,request *http.Request){
	per := persona.Persona{}

	db := conexion.Conexion()
	defer db.Close()

	id:= mux.Vars(request)["id"]
 


db.Find(&per,id)

if per.ID > 0 {
	db.Delete(per)
	handler.SendRespond(writer, http.StatusOK,[]byte(`{}`))
}else {
	handler.SendError(writer, http.StatusNotFound)
}

}


func Holamundo(writer http.ResponseWriter,request *http.Request){
	log.Println("Hola mundo")
}