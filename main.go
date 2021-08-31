package main

import (
	"log"
	"net/http"

	"github.com/GibranHL0/contact-api-go/app"
	"github.com/GibranHL0/contact-api-go/app/config"
)

func main() {
	app := app.New()
	mux := app.Server
	port := config.Get().AppPort

	defer app.Db.Close()

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
