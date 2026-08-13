package models

//import "image"

type Task struct {
	ID int
	Name string `json:"name"`
	PictureURL string
	Answer string `json:"answer,omitempty"`
	ParentID int `json:"parent_id"`
	TypeContent string `json:"type_content"`
}

type Chapter struct {
	ID int
	Name string
	ParentID int
}

type YandexConfig struct {
	YandexStaticEndpoint string `toml:"yandex_static_endpoint"`
	YandexRegion         string `toml:"yandex_region"`
	BacketName           string `toml:"backet_name"`
	AccessKeyID          string `toml:"access_key_ID"`
	SecretAccessKey      string `toml:"secret_access_key"`
}

