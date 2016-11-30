package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

const OutputDir = "output"

type BuildCommand struct {
}

func ConfigureBuildCommand(app *kingpin.Application) {
	cmd := &BuildCommand{}
	app.Command("build", "Build CV site").Action(cmd.run)
}

func (cmd *BuildCommand) run(c *kingpin.ParseContext) error {
	var userInfo userInfo
	userInfo.load()
	userInfo.copyTemplate()
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
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(b, model)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (info *userInfo) copyTemplate() {
	cfg := GetConfig()

	err := os.MkdirAll(OutputDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = copyDir(path.Join(TemplatesDir, cfg.Template.Name), OutputDir)
	if err != nil {
		log.Fatal(err)
	}
}

func (info *userInfo) render() {
	cfg := GetConfig()

	f, err := os.Create(path.Join(OutputDir, TemplateFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t, err := template.ParseFiles(path.Join(TemplatesDir, cfg.Template.Name, TemplateFileName))
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, info)
	if err != nil {
		log.Fatal(err)
	}
}

func copyDir(src, dst string) error {
	err := filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
		localPath, err := filepath.Rel(src, srcPath)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, localPath)

		if info.IsDir() {
			err = os.MkdirAll(dstPath, info.Mode())
			return err
		} else {
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			if err := dstFile.Chmod(info.Mode()); err != nil {
				return err
			}

			_, err = io.Copy(dstFile, srcFile)
			return err
		}
	})

	return err
}
