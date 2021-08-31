package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/GibranHL0/contact-api-go/app/config"
	"github.com/GibranHL0/contact-api-go/app/models"
	"github.com/GibranHL0/contact-api-go/app/repositories"
)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Open() error {
	uri := config.Get().DbUri

	pg, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		return err
	}

	p.db = pg

	p.db.AutoMigrate(models.Contact{})

	return nil
}

func (p *Postgres) Close() error {
	pg, err := p.db.DB()

	if err != nil {
		return err
	}

	return pg.Close()
}

func (p *Postgres) GetRepository() repositories.AbstractRepository {
	repo := repositories.GormRepository{}
	repo.GetDB(p.db)

	return &repo
}
