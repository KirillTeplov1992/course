package store

import (
	adminRep "course/internal/admin/storage"
	taskRep "course/internal/tasks/storage"
	"course/pkg/logging"
	"course/pkg/sqlite3"
)


type Store struct{
	TaskRepository *taskRep.TaskRepository
	AdminRepository *adminRep.AdminRepository
}

func NewStore(store *sqlite3.Store, logger *logging.Logger) *Store{
	return &Store{
		TaskRepository: taskRep.NewRepository(store, logger),
		AdminRepository: adminRep.NewRepository(store, logger),
	}
}