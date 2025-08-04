package notes

type repositoryContract interface {
	GetNotes() ([]Note, error)
	CreateNote(note Note) error
	EditNote(note editNoteParams, id string) (Note, error)
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

func (s service) createNote(noteRequest createNoteParams) error {
	err := noteRequest.Validate()
	if err != nil {
		return err
	}

	newNote := NewNote(noteRequest.Title, noteRequest.Link)
	err = s.repo.CreateNote(newNote)

	if err != nil {
		return err
	}
	return nil
}

func (s service) editNote(editRequest editNoteParams, id string) (note Note, err error) {
	note, err = s.repo.EditNote(editRequest, id)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}
