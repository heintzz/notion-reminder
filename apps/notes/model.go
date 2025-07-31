package notes

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Link           string `json:"link"`
	NextReminderAt string `json:"next_reminder_at"`
	CreatedAt      string `json:"created_at"`
}

func NewNote(id, title, link, nextReminderAt, createdAt string) Note {
	currentTime := time.Now()
	return Note{
		ID:             uuid.NewString(),
		Title:          title,
		Link:           link,
		NextReminderAt: currentTime.AddDate(0, 0, 1).Format(time.RFC3339),
		CreatedAt:      currentTime.Format(time.RFC3339),
	}
}
