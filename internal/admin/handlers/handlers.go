package handlers

import (
	"course/internal/handlers"
	"course/internal/store"
	"course/internal/tasks/models"
	"course/pkg/logging"
	"course/ui/templates"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	adminURL = "/admin"
	getDataURL = "/get_data"
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
	router.HandleFunc(getDataURL, h.getData)
}

func (h *handler) adminPage (w http.ResponseWriter, r *http.Request){
	title := "Добавление задания"
	chapters := h.repository.AdminRepository.GetAllChapters()

	c := templates.MainAdminPage(chapters)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}

func (h *handler) getData (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	
	task := models.Task{}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil{
		h.logger.Error("JSON не отправился", err)
	}

	fmt.Println(task.IsTask)

	//h.repository.AdminRepository.AddTask(task)

	

	// Отправляем ответ обратно
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "success", "message": "Данные успешно получены"}
	json.NewEncoder(w).Encode(response)
}