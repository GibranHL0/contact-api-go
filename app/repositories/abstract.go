package repositories

import "github.com/GibranHL0/contact-api-go/app/models"

// Base repository for all DBs.
type AbstractRepository interface {
	Add(model models.Model) (models.Model, error)
	GetById(id string, model models.Model) (models.Model, error)
	GetAll(models []models.Model) ([]models.Model, error)
	Update(id string, model models.Model, update models.Model) (models.Model, error)
	Delete(id string, model models.Model) error
}