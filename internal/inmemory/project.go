package inmemory

import "go.lorenzomilicia.com/go-master/http-server/internal/model"

type InMemoryProjectRepository struct {
	Projects []model.Project
}

func (r *InMemoryProjectRepository) GetAll() ([]model.Project, error) {
	return r.Projects, nil
}

func (r *InMemoryProjectRepository) GetByUid(uid string) (*model.Project, error) {
	var project *model.Project
	for idx, p := range r.Projects {
		if p.UID == uid {
			project = &r.Projects[idx]
			break
		}
	}
	return project, nil
}

func (r *InMemoryProjectRepository) Save(project model.Project) error {
	r.Projects = append(r.Projects, project)
	return nil
}
