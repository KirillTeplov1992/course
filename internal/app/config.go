package app

import "course/pkg/sqlite3"

type Config struct{
	BindAddr string
	Store *sqlite3.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: "5010",
		Store: sqlite3.NewConfig(),
	}
}