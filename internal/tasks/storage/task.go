package storage

import (
	"course/internal/tasks/models"
	"course/pkg/logging"
	"course/pkg/sqlite3"
)


type TaskRepository struct{
	store *sqlite3.Store
	logger *logging.Logger
}

func NewRepository(store *sqlite3.Store, logger *logging.Logger) *TaskRepository{
	return &TaskRepository{
		store: store,
		logger: logger,
	}
}

func (tr *TaskRepository) GetContents() []*models.Chapter {
	stmt := `
	select 
		id,
		task
	from
		tasks
	WHERE
 		parent_id is NULL`

	res, err := tr.store.DB.Query(stmt)
	if err != nil {
		tr.logger.Fatal(err)
	}

	var ChapterList []*models.Chapter

	for res.Next(){
		chapter := &models.Chapter{}
		err = res.Scan(&chapter.ID, &chapter.Name)
		if err != nil{
			tr.logger.Fatal(err)
		}

		ChapterList = append(ChapterList, chapter)
	}

	return ChapterList
}

func (tr *TaskRepository) GetTasks(id int) []*models.Chapter{
	stmt := `
	select 
		id,
		task
	from
		tasks
	WHERE
 		parent_id = ?`

	res, err := tr.store.DB.Query(stmt, id)
	if err != nil{
		tr.logger.Fatal(err)
	}

	var taskList []*models.Chapter

	for res.Next(){
		task := &models.Chapter{}
		err = res.Scan(&task.ID, &task.Name)
		if err != nil{
			tr.logger.Fatal(err)
		}

		taskList = append(taskList, task)
	}

	return taskList
}

func (tr *TaskRepository) GetNameChapterById(id int) string{
	stmt := `select task from tasks where id = ?`

	row := tr.store.DB.QueryRow(stmt, id)

	var title string

	if err := row.Scan(&title); err != nil{
		tr.logger.Fatal(err)
	}

	return title
}