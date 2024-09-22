package db

type DatabaseConfig struct {
	Port int
	IP   string

	Login    string
	Password string

	Database string
}

func CreateConfig() *DatabaseConfig {
	return &DatabaseConfig{}
}
