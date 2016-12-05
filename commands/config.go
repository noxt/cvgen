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
	RepoURL string `yaml:"repo_url"`
	Path    string
	Files   []string
}

var configInstance *config
var once sync.Once

// Get current config
func GetConfig() *config {
	once.Do(func() {
		configInstance = &config{}

		b, err := ioutil.ReadFile(configFileName)
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
