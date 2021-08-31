package database

import "github.com/GibranHL0/contact-api-go/app/repositories"

type Database interface {
	Open() error
	Close() error
	GetRepository() repositories.AbstractRepository
}
