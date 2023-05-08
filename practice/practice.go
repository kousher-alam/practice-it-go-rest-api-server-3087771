package practice

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! from practice!\n")
}

func RunServer(port string) {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server is running on localhost", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
