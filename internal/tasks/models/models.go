package models

import "image"

type Task struct {
	ID int
	Name string
	Picture image.Image
	Answer string
	ParentID int
}

type Chapter struct {
	ID int
	Name string
	ParentID int
}