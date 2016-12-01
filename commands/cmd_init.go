package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	DefaultTemplateRepo     = "https://github.com/noxt/cvgen-templates"
	DefaultTemplatePath     = "orbit"
	DefaultTemplateFileName = "index.html"
)

type InitCommand struct {
}

func ConfigureInitCommand(app *kingpin.Application) {
	cmd := &InitCommand{}
	app.Command("init", "Initialize project structure").Action(cmd.run)
}

func (cmd *InitCommand) run(c *kingpin.ParseContext) error {
	var parsingMap = map[string]interface{}{
		ConfigFileName: config{
			Template: templateRepo{
				RepoURL: DefaultTemplateRepo,
				Path:    DefaultTemplatePath,
				Files:   []string{DefaultTemplateFileName},
			},
		},
		AboutMeFileName: me{
			Languages: []language{
				language{},
			},
		},
		EducationFileName: []education{
			education{},
		},
		OrganizationsFileName: []organization{
			organization{
				Projects: []project{
					project{},
				},
			},
		},
		ProjectsFileName: []project{
			project{},
		},
		SkillsFileName: []skill{
			skill{},
		},
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

	return nil
}
