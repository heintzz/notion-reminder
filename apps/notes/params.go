package notes

import (
	"heintzz/notion-reminder/internal/helper"
)

type createNoteParams struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type editNoteParams struct {
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
}

func (req createNoteParams) Validate() error {
	if req.Title == "" {
		return helper.ErrTitleRequired
	}
	if req.Link == "" {
		return helper.ErrLinkRequired
	}
	return nil
}
