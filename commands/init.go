package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	DefaultTemplateRepo = "https://github.com/noxt/cvgen-templates"
	DefaultTemplateName = "orbit"
)

type InitCommand struct {
}

func (cmd *InitCommand) run(c *kingpin.ParseContext) error {
	var parsingMap = map[string]interface{}{
		ConfigFileName: config{
			Templates: templatesRepo{
				Repo: DefaultTemplateRepo,
				Name: DefaultTemplateName,
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
		CheckIfError(err)

		err = ioutil.WriteFile(file, b, os.ModePerm)
		CheckIfError(err)
	}

	return nil
}

func ConfigureInitCommand(app *kingpin.Application) {
	cmd := &InitCommand{}
	app.Command("init", "Initialize project structure").Action(cmd.run)
}
