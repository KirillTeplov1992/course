package handlers

import (
	"course/internal/handlers"
	"course/internal/store"
	"course/pkg/logging"
	"course/ui/templates"
	"net/http"
	"strconv"
)

const (
	homeURL =  "/"
	taskURL = "/task"
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

func (h *handler) Register(router *http.ServeMux){
	router.HandleFunc(homeURL, h.home)
	router.HandleFunc(taskURL, h.getTask)

	//подключаю CSS стили
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.Handle("/static/", http.StripPrefix("/static", fileServer))
}

func (h *handler) home(w http.ResponseWriter, r *http.Request){
	title := "Домашняя страница"
	chapters := h.repository.TaskRepository.GetContents()

	c := templates.Home(chapters)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}

func (h *handler) getTask(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	tasks := h.repository.TaskRepository.GetTasks(id)
	title := h.repository.TaskRepository.GetNameChapterById(id)

	c := templates.Home(tasks)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}
