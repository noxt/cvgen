package commands

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type config struct {
	Templates templatesRepo
}

type templatesRepo struct {
	Repo string
	Name string
}

var configInstance *config
var once sync.Once

func GetConfig() *config {
	once.Do(func() {
		configInstance = &config{}

		b, err := ioutil.ReadFile(ConfigFileName)
		CheckIfError(err)

		err = yaml.Unmarshal(b, configInstance)
		CheckIfError(err)
	})
	return configInstance
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
