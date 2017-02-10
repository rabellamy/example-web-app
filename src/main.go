package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/going/defaults"
	"github.com/rabellamy/example-web-app/src/actions"
)

func main() {
	port := defaults.String(os.Getenv("PORT"), "3000")
	log.Printf("Starting example-web-app on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
