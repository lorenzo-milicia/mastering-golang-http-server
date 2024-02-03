package model

import (
	"time"
)

type Project struct {
	UID            string
	Name           string
	DateOfCreation time.Time
}

type ProjectRepository interface {
	GetAll() ([]Project, error)
	GetByUid(uid string) (*Project, error)
	Save(Project) error
}
