package services

import (
	"github.com/GibranHL0/contact-api-go/app/models"
	"github.com/GibranHL0/contact-api-go/app/repositories"
)

func GetContacts(repo repositories.AbstractRepository, contacts []models.Model) ([]models.Contact, error) {

	contacts, err := repo.GetAll(contacts)

	if err != nil {
		return nil, err
	}
	
	return contacts, nil
}