package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")
	router := mux.NewRouter()
	const port string = "8000"
}
