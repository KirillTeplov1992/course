package yandex

import (
	"course/internal/tasks/models"
	"log"

	"github.com/BurntSushi/toml"
)

func NewConfig() *models.YandexConfig{
	var Config models.YandexConfig

	_, err := toml.DecodeFile("/home/kirill/go/couse/configs/yandex_config.toml", &Config)
	if err != nil{
	log.Println("Config не парсится")
	}

return &Config
}