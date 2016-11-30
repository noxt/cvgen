package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"text/template"
)

const OutputDir = "output"

type BuildCommand struct {
}

func (cmd *BuildCommand) run(c *kingpin.ParseContext) error {
	var userInfo userInfo
	userInfo.load()
	userInfo.render()

	return nil
}

func (info *userInfo) load() {
	var parsingMap = map[string]interface{}{
		AboutMeFileName:       &info.AboutMe,
		EducationFileName:     &info.Educations,
		OrganizationsFileName: &info.Organizations,
		ProjectsFileName:      &info.Projects,
		SkillsFileName:        &info.Skills,
	}

	for file, model := range parsingMap {
		b, err := ioutil.ReadFile(filepath.Join("./", file))
		CheckIfError(err)

		err = yaml.Unmarshal(b, model)
		CheckIfError(err)
	}
}

func (info *userInfo) render() {
	cfg := GetConfig()
	t, err := template.ParseFiles(path.Join(TemplatesDir, cfg.Template.Name, TemplateFileName))
	CheckIfError(err)

	err = os.MkdirAll(OutputDir, os.ModePerm)
	CheckIfError(err)

	// TODO: Исправить копирование
	c := exec.Command("/bin/sh", "-c", "cp -a "+path.Join(TemplatesDir, cfg.Template.Name, "*")+" "+OutputDir)
	err = c.Run()
	CheckIfError(err)

	f, err := os.Create(path.Join(OutputDir, TemplateFileName))
	CheckIfError(err)
	defer f.Close()

	err = t.Execute(f, info)
	CheckIfError(err)
}

func ConfigureBuildCommand(app *kingpin.Application) {
	cmd := &BuildCommand{}
	app.Command("build", "Build CV site").Action(cmd.run)
}
