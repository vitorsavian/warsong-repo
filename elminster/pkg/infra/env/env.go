package environment

import (
	libenv "github.com/Netflix/go-env"
	"github.com/sirupsen/logrus"
)

type Env struct {
}

func GetEnv() *Env {
	var env *Env
	_, err := libenv.UnmarshalFromEnviron(env)
	if err != nil {
		logrus.Fatalf("Env deu pal")
	}
	return env
}
