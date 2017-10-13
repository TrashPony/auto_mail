package main

import (
	//
	"github.com/gorilla/mux"
	"./webSocket"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", webSocket.HandleConnections) // если браузер запрашивает соеденение на /ws то инициализируется переход на вебсокеты
	go webSocket.ReposeSender()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("http server started on :3000")
	http.ListenAndServe(":3000", router)
}
