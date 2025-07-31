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

func (h handler) createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req createNoteParams

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := map[string]interface{}{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "Invalid request payload",
		}
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	err = h.svc.createNote(req)
	if err != nil {
		resp := map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"error":   "Internal Server Error",
			"message": "Failed to create note",
		}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	w.Write([]byte("mantap"))
}
