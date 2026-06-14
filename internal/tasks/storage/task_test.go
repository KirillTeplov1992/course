package storage

/*import (
	"course/pkg/logging"
	"course/pkg/sqlite3"
	"testing"
)



func TestGetContents(t *testing.T){
	config := sqlite3.NewConfig()
	s := sqlite3.NewStore(config)
	logget := logging.GetLogger()
	rep := NewRepository(s, &logget)

	/*expected := []*models.Chapter{&models.Chapter{
		ID: 1,
		Name: "Тесты",
	},
	}

	_, err := rep.GetContents()

	if err != nil{
		t.Error("Жопа")
	}
}*/