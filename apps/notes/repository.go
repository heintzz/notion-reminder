package notes

import (
	"encoding/json"
	"errors"
	"heintzz/notion-reminder/internal/helper"
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

func (r repository) GetNoteByID(id string) (note Note, err error) {
	notes, err := r.GetNotes()
	if err != nil {
		return Note{}, helper.ErrNoteNotFound
	}

	for _, note := range notes {
		if note.ID == id {
			n := note
			return n, nil
		}
	}

	return
}

func (r repository) CreateNote(newNote Note) (err error) {
	notes, err := r.GetNotes()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	notes = append(notes, newNote)

	jsonData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filepath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) EditNote(editRequest editNoteParams, id string) (note Note, err error) {
	notes, err := r.GetNotes()
	if err != nil && !os.IsNotExist(err) {
		return Note{}, err
	}

	var newNote Note
	var isExist bool = false

	for i, n := range notes {
		if n.ID == id {
			if editRequest.Link != "" {
				notes[i].Link = editRequest.Link
			}
			if editRequest.Title != "" {
				notes[i].Title = editRequest.Title
			}
			newNote = notes[i]
			isExist = true
			break
		}
	}

	if !isExist {
		return Note{}, errors.New("note not found")
	}

	jsonData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return Note{}, err
	}

	err = os.WriteFile(r.filepath, jsonData, 0644)
	if err != nil {
		return Note{}, err
	}

	return newNote, nil
}
