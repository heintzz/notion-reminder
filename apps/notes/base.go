package notes

import "github.com/go-chi/chi/v5"

func Run(router chi.Router) {
	repo := newRepository()
	svc := newService(repo)
	handler := newHandler(svc)

	router.Route("/v1/notes", func(r chi.Router) {
		r.Get("/", handler.getNotesHandler)
		r.Post("/", handler.createNoteHandler)
		r.Put("/{id}", handler.updateNoteHandler)
	})
}
