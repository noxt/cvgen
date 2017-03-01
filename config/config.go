package config

import (
	"fmt"
	"github.com/noxt/cvgen/constants"
	"github.com/noxt/cvgen/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var configInstance *models.Config
var once sync.Once

func NewConfig() *models.Config {
	return &models.Config{
		Template: models.TemplateRepo{
			RepoURL: "https://github.com/noxt/cvgen-templates",
			Path:    "orbit",
			Files:   []string{"index.html"},
		},
		OutputDir: "output",
	}
}

// GetConfig return loaded from file config
func GetConfig() *models.Config {
	once.Do(func() {
		configInstance = NewConfig()

		b, err := ioutil.ReadFile(constants.ConfigFileName)
		if err != nil {
			return
		}

		err = yaml.Unmarshal(b, configInstance)
		if err != nil {
			log.Fatal(fmt.Errorf("config: decode YAML: %v", err))
		}
	})

	return configInstance
}
