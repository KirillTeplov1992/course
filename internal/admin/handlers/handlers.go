package handlers

import (
	"course/internal/admin/yandex"
	"course/internal/handlers"
	"course/internal/store"
	"course/internal/tasks/models"
	"course/pkg/logging"
	"course/ui/templates"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	adminURL = "/admin"
	getDataURL = "/get_data"
	addChapterURL = "/admin/add_chapter"
	saveChapterURL = "/admin/save_chapter"
	addTaskURL = "/admin/add_task"
	addArticleURL = "/admin/add_article"
	addTaskWithPictureURL = "/admin/add_task_with_picture"
	getTaskWithPictureURL = "/get_task_with_picture"
)

var _ handlers.Handler = &handler{}

type handler struct{
	logger logging.Logger
	repository store.Store
	yandexClient yandex.S3Client
}

func NewHandler(repository store.Store, logger logging.Logger) handlers.Handler{
	return &handler{
		logger: logger,
		repository: repository,
		yandexClient: *yandex.NewS3Client(),
	}
}

func (h *handler) Register (router *http.ServeMux){
	router.HandleFunc(adminURL, h.adminPage)
	router.HandleFunc(getDataURL, h.getData)
	router.HandleFunc(addChapterURL, h.addChapter)
	router.HandleFunc(saveChapterURL, h.saveChapter)
	router.HandleFunc(addArticleURL, h.addArticle)
	router.HandleFunc(addTaskURL, h.addTask)
	router.HandleFunc(addTaskWithPictureURL, h.addTaskWithPicture)
	router.HandleFunc(getTaskWithPictureURL, h.getTaskWithPicture)
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
		PictureURL: "",
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

func(h *handler) getTaskWithPicture(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	// 1. Парсим multipart форму (максимум 10 МБ на запрос)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Получаем текстовые поля через r.FormValue
	chapterID, err := strconv.Atoi(r.FormValue("chapter"))
	if err != nil {
		h.logger.Fatal(err)
	}
	taskCode := r.FormValue("task")
	answerCode := r.FormValue("answer")
	typeContent := r.FormValue("type_content")
	
	
	// (Здесь ваша логика сохранения текстов в базу данных)
	fmt.Printf("Глава: %d, Получена задача: %s, Ответ: %s, Тип: %s\n",chapterID, taskCode, answerCode, typeContent)

	// 3. Обрабатываем загруженный файл картинки
	file, header, err := r.FormFile("image")
	if err == nil {
		// Ошибка nil означает, что файл успешно прикреплен к запросу
		defer file.Close()
	}

	uniqueFileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), header.Filename)
	fileURL, err := h.yandexClient.DownLoadFile(uniqueFileName, file)
	if err != nil{
		h.logger.Printf("Ошибка загрузки в S3 %v", err)
		http.Error(w, "Не удалось загрузить файл в облако", http.StatusInternalServerError)
	}

	fmt.Println(fileURL)

	task := models.Task{
		Name: taskCode,
		PictureURL: fileURL,
		Answer: answerCode,
		ParentID: chapterID,
		TypeContent: typeContent,
	}

	h.repository.AdminRepository.AddTask(task)					

	// 4. Отправляем успешный ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "success", "message": "Данные успешно получены"}
	json.NewEncoder(w).Encode(response)
}