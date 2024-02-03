package project

import "go.lorenzomilicia.com/go-master/http-server/internal/model"

type Service struct {
	Repository model.ProjectRepository
}

func (s *Service) GetAll() ([]model.Project, error) {
	return s.Repository.GetAll()
}

func (s *Service) GetByUid(uid string) (*model.Project, error) {
	return s.Repository.GetByUid(uid)
}

func (s *Service) Save(project model.Project) error {
	return s.Repository.Save(project)
}
