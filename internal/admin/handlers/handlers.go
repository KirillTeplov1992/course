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
	"strconv"
)

const (
	adminURL = "/admin"
	getDataURL = "/get_data"
	addChapterURL = "/admin/add_chapter"
	saveShapterURL = "/admin/save_chapter"
	addTaskURL = "/admin/add_task"
	addArticleURL = "/admin/add_article"
	addTaskWithPictureURL = "/admin/add_task_with_picture"
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
	router.HandleFunc(addChapterURL, h.addChapter)
	router.HandleFunc(saveShapterURL, h.saveChapter)
	router.HandleFunc(addArticleURL, h.addArticle)
	router.HandleFunc(addTaskURL, h.addTask)
	router.HandleFunc(addTaskWithPictureURL, h.addTaskWithPicture)
}

func (h *handler) adminPage (w http.ResponseWriter, r *http.Request){
	title := "Страница админа"

	c := templates.MainAdminPage()
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}

func (h *handler) addTask (w http.ResponseWriter, r *http.Request){
	title := "Добавить задание"

	chapters := h.repository.AdminRepository.GetAllChapters()

	c := templates.AddTask(chapters)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}

func (h *handler) addChapter (w http.ResponseWriter, r *http.Request){
	title := "Добавить раздел"

	chapters := h.repository.AdminRepository.GetAllChapters()

	c := templates.AddChapters(chapters)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}

func (h *handler) addTaskWithPicture (w http.ResponseWriter, r *http.Request){
	title := "Добавить задание с картинкой"

	chapters := h.repository.AdminRepository.GetAllChapters()

	c := templates.AddTaskWithPictere(chapters)
	if err := templates.Layout(c, title).Render(r.Context(), w); err != nil{
		h.logger.Fatal(err)
	}
}

func (h *handler) saveChapter(w http.ResponseWriter, r *http.Request){
	parent_id, err := strconv.Atoi(r.FormValue("chapter_id"))
	if err != nil {
		h.logger.Fatal(err)
	}
	chapter := r.FormValue("chapter")

	task := models.Task{
		Name: chapter,
		ParentID: parent_id,
		TypeContent: "title",
	}

	h.repository.AdminRepository.AddTask(task)
}

func (h *handler) addArticle (w http.ResponseWriter, r *http.Request){
	title := "Добавить статью"

	chapters := h.repository.AdminRepository.GetAllChapters()

	c := templates.AddArticle(chapters)
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

	fmt.Println(task.TypeContent)

	h.repository.AdminRepository.AddTask(task)	

	// Отправляем ответ обратно
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "success", "message": "Данные успешно получены"}
	json.NewEncoder(w).Encode(response)
}