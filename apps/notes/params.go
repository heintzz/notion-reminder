package notes

import "errors"

type createNoteParams struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type createNoteResponse struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Link           string `json:"link"`
	NextReminderAt string `json:"next_reminder_at"`
	CreatedAt      string `json:"created_at"`
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
