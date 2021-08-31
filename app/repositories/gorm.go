package repositories

import (
	"gorm.io/gorm"
	
	"github.com/GibranHL0/contact-api-go/app/models"
)

type GormRepository struct {
	db *gorm.DB
}

func (gr *GormRepository) GetDB(db *gorm.DB) {
	gr.db = db
}

func (gr *GormRepository) Add(model models.Model) (models.Model, error) {
	query := gr.db.Create(&model)

	if query.Error != nil {
		return nil, query.Error
	}

	return model, nil
}

func (gr *GormRepository) GetById(id string, model models.Model) (models.Model, error) {
	query := gr.db.First(&model, id)

	if query.Error != nil {
		return nil, query.Error
	}

	return model, nil
}

func (gr *GormRepository) GetAll(models []models.Model) ([]models.Model, error) {
	query := gr.db.Find(&models)

	if query.Error != nil {
		return nil, query.Error
	}

	return models, nil
}

func (gr *GormRepository) Update(id string, model models.Model, update models.Model) (models.Model, error) {
	query := gr.db.Model(model).Where("id = ?", id).Updates(update)

	if query.Error != nil{
		return nil, query.Error
	}

	return update, nil
}

func (gr *GormRepository) Delete(id string, model models.Model) error {
	query := gr.db.Where("id = ?", id).Delete(model)

	if query.Error != nil {
		return query.Error
	}

	return nil
}