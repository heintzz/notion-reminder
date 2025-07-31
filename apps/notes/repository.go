package notes

import (
	"encoding/json"
	"os"
)

var filepath = "internal/data/notes.json"

type repository struct {
	filepath string
}

func newRepository() repository {
	return repository{
		filepath: filepath,
	}
}

func (r repository) GetNotes() (notes []Note, err error) {
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
