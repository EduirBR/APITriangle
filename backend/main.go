package main

import (
	"fmt"
	"log"
	"net/http"
	"tri/cors"
	"tri/handlers"

	"github.com/gorilla/mux"
)

func main() {

mux:=mux.NewRouter()
cors.EnableCORS(mux)

mux.HandleFunc("/triangle", handlers.TriangleHandle).Methods(http.MethodPost)

server := &http.Server{
	Addr: "localhost:8000",
	Handler: mux,
}
fmt.Println("Server Runing at:")
fmt.Println("localhost:8000")
log.Fatal(server.ListenAndServe())

}



