package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO: Continue with another endpoint here
// https://youtu.be/W5b64DXeP0o

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
