package notes

type repositoryContract interface {
	GetNotes() ([]Note, error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) getNotes() ([]Note, error) {
	notes, err := s.repo.GetNotes()
	if err != nil {
		return nil, err
	}

	return notes, nil
}
