package notes

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) getNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := h.svc.getNotes()
	if err != nil {
		http.Error(w, "Failed to get notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(notes) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	jsonResponse, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
