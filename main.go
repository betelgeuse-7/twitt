package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	//!  Don't forget pagination
	routes()

	log.Println(fmt.Sprintf("Starting server at localhost%s\n", PORT))
	log.Fatalln(http.ListenAndServe(PORT, r))
}
