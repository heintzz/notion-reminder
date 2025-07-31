package notes

import "errors"

type createNoteParams struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func (req createNoteParams) Validate() error {
	if req.Title == "" {
		return errors.New("note title is required")
	}
	if req.Link == "" {
		return errors.New("note link is required")
	}
	return nil
}
