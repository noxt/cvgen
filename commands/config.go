package commands

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type config struct {
	Template templateRepo
}

type templateRepo struct {
	RepoURL  string `yaml:"repo_url"`
	Path     string
	FileName string `yaml:"file_name"`
}

var configInstance *config
var once sync.Once

func GetConfig() *config {
	once.Do(func() {
		configInstance = &config{}

		b, err := ioutil.ReadFile(ConfigFileName)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(b, configInstance)
		if err != nil {
			log.Fatal(err)
		}
	})
	return configInstance
}
