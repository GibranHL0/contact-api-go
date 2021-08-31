package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GibranHL0/contact-api-go/app/models"
	"github.com/GibranHL0/contact-api-go/app/repositories"
	"github.com/GibranHL0/contact-api-go/app/services"
)

func GetContacts(repo repositories.AbstractRepository,w http.ResponseWriter, r *http.Request) {
	con := []models.Contact{}
	contacts, err := services.GetContacts(repo, con)

	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}

func GetContactById(w http.ResponseWriter, r *http.Request) {}

func UploadContact(w http.ResponseWriter, r *http.Request) {}

func UpdateContact(w http.ResponseWriter, r *http.Request) {}

func DeleteContact(w http.ResponseWriter, r *http.Request) {}
