package commands

import (
	"fmt"
	"github.com/noxt/cvgen/config"
	"github.com/noxt/cvgen/constants"
	"github.com/noxt/cvgen/models"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// ConfigureInitCommand setup "init" command
func ConfigureInitCommand(app *kingpin.Application) {
	app.Command("init", "Initialize YAML files into current directory").Action(runInitCommand)
}

func runInitCommand(*kingpin.ParseContext) error {
	generateYAMLFiles()
	return nil
}

func generateYAMLFiles() {
	log.Println("init: generate YAML files: start")

	var parsingMap = map[string]interface{}{
		constants.ConfigFileName:        config.NewConfig(),
		constants.AboutMeFileName:       models.Me{Languages: []models.Language{{}}},
		constants.EducationFileName:     []models.Education{{}},
		constants.OrganizationsFileName: []models.Organization{{Projects: []models.Project{{}}}},
		constants.ProjectsFileName:      []models.Project{{}},
		constants.SkillsFileName:        []models.Skill{{}},
	}

	for file, model := range parsingMap {
		b, err := yaml.Marshal(model)
		if err != nil {
			log.Fatal(fmt.Errorf("init: encoding YAML: %v", err))
		}

		err = ioutil.WriteFile(file, b, os.ModePerm)
		if err != nil {
			log.Fatal(fmt.Errorf("init: save YAML file: %v", err))
		}
	}

	log.Println("init: generate YAML files: finish")
}
