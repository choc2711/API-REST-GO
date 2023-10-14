package handler

import "net/http"

func SendRespond(writer http.ResponseWriter, status int, data []byte){
	writer.Header().Set("Content-Type", "aplication/json")
	writer.WriteHeader(status)
	writer.Write(data)
}

func SendError(writer http.ResponseWriter, status int)  {
	data := []byte(`{}`)
	writer.Header().Set("Content_Type","application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}