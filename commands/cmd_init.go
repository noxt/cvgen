package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultTemplateRepo     = "https://github.com/noxt/cvgen-templates"
	defaultTemplatePath     = "orbit"
	defaultTemplateFileName = "index.html"
)

// ConfigureInitCommand setup "init" command
func ConfigureInitCommand(app *kingpin.Application) {
	app.Command("init", "Initialize YAML files into current directory").Action(runInitCommand)
}

func runInitCommand(*kingpin.ParseContext) error {
	var parsingMap = map[string]interface{}{
		configFileName: config{
			Template: templateRepo{
				RepoURL: defaultTemplateRepo,
				Path:    defaultTemplatePath,
				Files:   []string{defaultTemplateFileName},
			},
		},
		aboutMeFileName:       me{Languages: []language{{}}},
		educationFileName:     []education{{}},
		organizationsFileName: []organization{{Projects: []project{{}}}},
		projectsFileName:      []project{{}},
		skillsFileName:        []skill{{}},
	}

	for file, model := range parsingMap {
		b, err := yaml.Marshal(model)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(file, b, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Successful initialization")

	return nil
}
