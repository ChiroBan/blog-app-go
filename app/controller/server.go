package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/pilinux/gorest/controller"
)

var router *mux.Router

func initHandlers() {
	
	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/post/{id}", controller.GetPost).Methods("GET")
	router.HandleFunc("/api/post/new", controller.CreatePost).Methods("POST")
}

func Start() {
	router = mux.NewRouter()
	
	initHandlers()
	fmt.Printf("router initialized and listening on 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}