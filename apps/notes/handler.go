package notes

import (
	"encoding/json"
	"heintzz/notion-reminder/internal/helper"
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
	var request createNoteParams

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		resp := map[string]interface{}{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "Invalid request payload",
		}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	err = h.svc.createNote(request)
	if err != nil {
		errors, ok := helper.ErrorMapping[err.Error()]
		if !ok {
			errors = helper.ErrorGeneral
		}
		resp := map[string]interface{}{
			"status":  errors.HttpCode,
			"error":   errors.Error,
			"message": errors.Message,
		}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Note created successfully",
	})
}

func (h handler) updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var request editNoteParams
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		resp := map[string]interface{}{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "Invalid request payload",
		}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	updatedNote, err := h.svc.editNote(request, r.PathValue("id"))
	if err != nil {
		resp := map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"error":   "Internal Server Error",
			"message": err.Error(),
		}
		jsonResp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Note edited successfully",
		"data":    updatedNote,
	})
}
