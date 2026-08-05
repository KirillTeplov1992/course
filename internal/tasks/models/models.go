package models

//import "image"

type Task struct {
	ID int
	Name string `json:"name"`
	//Picture image.Image
	Answer string `json:"answer,omitempty"`
	ParentID int `json:"parent_id"`
	TypeContent string `json:"type_content"`
}

type Chapter struct {
	ID int
	Name string
	ParentID int
}