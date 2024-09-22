package environment

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func SetEnv(env string) {
	if env != "prd" {
		err := godotenv.Load()
		if err != nil {
			logrus.Fatal("Error loading .env file")
		}
	}
}
