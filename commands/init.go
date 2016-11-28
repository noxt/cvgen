package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
		AboutMeFileName: me{},
		EducationFileName: []education{
			education{},
		},
		OrganizationsFileName: []organization{
			organization{},
		},
		ProjectsFileName: []project{
			project{},
		},
	}

	for file, model := range parsingMap {
		b, err := yaml.Marshal(model)
		CheckIfError(err)

		err = ioutil.WriteFile(file, b, 0755)
		CheckIfError(err)
	}

	return nil
}

func ConfigureInitCommand(app *kingpin.Application) {
	cmd := &InitCommand{}
	app.Command("init", "Initialize project structure").Action(cmd.run)
}
