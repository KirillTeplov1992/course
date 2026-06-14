package handlers

import (
	"course/internal/handlers"
	"course/internal/store"
	"course/pkg/logging"
	"course/ui/templates"
	"net/http"
)

const (
	adminURL = "/admin"
)

var _ handlers.Handler = &handler{}

type handler struct{
	logger logging.Logger
	repository store.Store
}

func NewHandler(repository store.Store, logger logging.Logger) handlers.Handler{
	return &handler{
		logger: logger,
		repository: repository,
	}
}

func (h *handler) Register (router *http.ServeMux){
	router.HandleFunc(adminURL, h.adminPage)
}

func (h *handler) adminPage (w http.ResponseWriter, r *http.Request){
	title := "Добавление задания"
	chapters := h.repository.AdminRepository.GetAllChapters()

	c := templates.MainAdminPage(chapters)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}
