package models

//import "image"

type Task struct {
	ID int
	Name string `json:"name"`
	//Picture image.Image
	Answer string `json:"answer,omitempty"`
	ParentID int `json:"parent_id"`
	IsTask bool `json:"is_task"`
}

type Chapter struct {
	ID int
	Name string
	ParentID int
}