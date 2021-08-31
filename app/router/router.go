package router

import (
	"net/http"

	"github.com/GibranHL0/contact-api-go/app/database"
	"github.com/GibranHL0/contact-api-go/app/handler"
)

func ContactsRouter(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := db.GetRepository()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case "GET":
			handler.GetContacts(repo, w, r)
		case "POST":
			handler.UploadContact(w, r)
		case "PUT":
			handler.UpdateContact(w, r)
		case "DELETE":
			handler.DeleteContact(w, r)
		}

	}
}
