package storage

import (
	"course/internal/tasks/models"
	"course/pkg/logging"
	"course/pkg/sqlite3"
)

type AdminRepository struct{
	store *sqlite3.Store
	logger *logging.Logger
}

func NewRepository(store *sqlite3.Store, logger *logging.Logger) *AdminRepository{
	return &AdminRepository{
		store: store,
		logger: logger,
	}
}

func (ar *AdminRepository) GetAllChapters() []*models.Chapter{
	stmt := `select id, task from tasks where type_content = "title"`

	res, err := ar.store.DB.Query(stmt)
	if err != nil{
		ar.logger.Fatal(err)
	}

	var Chapters []*models.Chapter

	for res.Next(){
		chapter := &models.Chapter{}

		if err = res.Scan(&chapter.ID, &chapter.Name); err != nil{
			ar.logger.Fatal(err)
		}

		Chapters = append(Chapters, chapter)
	}
	
	return Chapters
}

func (ar *AdminRepository) AddTask(task models.Task){
	stmt := `
	INSERT INTO tasks (
		task,
		answer,
		parent_id,
		type_content	
	)
	VALUES (
		?,
		?,
		?,
		?
	)`

	_, err := ar.store.DB.Exec(stmt,
								task.Name,
								task.Answer,
								task.ParentID,
								task.TypeContent)
	if err != nil{
		ar.logger.Fatal(err)
	}
}
