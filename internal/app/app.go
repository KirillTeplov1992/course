package app

import (
	adminHandler "course/internal/admin/handlers"
	"course/internal/store"
	taskHandler "course/internal/tasks/handlers"
	"course/pkg/logging"
	"course/pkg/sqlite3"
	"net/http"
)


type Application struct{
	router *http.ServeMux
	store *store.Store
	logger logging.Logger
	config *Config
}

func New() *Application{
	return &Application{
		router: http.NewServeMux(),
		logger: logging.GetLogger(),
		config: NewConfig(),
	}
}

func (app *Application) Start() error{
	app.logger.Info("Запуск веб-сервера на http://127.0.0.1:4000")

	if err := app.configureStore(); err != nil {
		return err
	}

	Handler := taskHandler.NewHandler(*app.store, app.logger)
	Handler.Register(app.router)

	adminHandler := adminHandler.NewHandler(*app.store, app.logger)
	adminHandler.Register(app.router)

	return http.ListenAndServe(":4000", app.router)
}

func (app *Application) configureStore() error{
	st := sqlite3.NewStore(app.config.Store)
	if err := st.Open(); err != nil{
		return err
	}

	app.store = store.NewStore(st, &app.logger)

	return nil
}