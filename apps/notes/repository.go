package notes

import (
	"encoding/json"
	"errors"
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

func (r repository) CreateNote(newNote Note) (err error) {
	var notes []Note
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &notes)
	if err != nil {
		return err
	}

	notes = append(notes, newNote)

	_, err = os.Stat(r.filepath)
	if os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return errors.New("error creating new file")
	}

	jsonData, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
