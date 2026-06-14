package sqlite3

type Config struct{
	DataBasePath string
}

func NewConfig() *Config{
	return &Config{
		DataBasePath: "/home/kirill/python/stocks/course.db",
	}
}