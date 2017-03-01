package commands

import (
	"fmt"
	"github.com/noxt/cvgen/config"
	"github.com/noxt/cvgen/constants"
	"github.com/noxt/cvgen/models"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// ConfigureBuildCommand setup "build" command
func ConfigureBuildCommand(app *kingpin.Application) {
	app.Command("build", "Build CV site from current directory").Action(runBuildCommand)
}

func runBuildCommand(*kingpin.ParseContext) error {
	userInfo := models.UserInfo{}

	parseYAMLFiles(&userInfo)
	copyTemplate()
	renderTemplate(&userInfo)

	return nil
}

func parseYAMLFiles(info *models.UserInfo) {
	log.Println("build: parsing YAML files: start")

	var parsingMap = map[string]interface{}{
		constants.AboutMeFileName:       &info.AboutMe,
		constants.EducationFileName:     &info.Educations,
		constants.OrganizationsFileName: &info.Organizations,
		constants.ProjectsFileName:      &info.Projects,
		constants.SkillsFileName:        &info.Skills,
	}

	for file, model := range parsingMap {
		b, err := ioutil.ReadFile(filepath.Join("./", file))
		if err != nil {
			break
		}

		err = yaml.Unmarshal(b, model)
		if err != nil {
			log.Fatal(fmt.Errorf("build: decoding YAML: %v", err))
		}
	}

	log.Println("build: parsing YAML files: finish")
}

func copyTemplate() {
	log.Println("build: copy template: start")

	cfg := config.GetConfig()

	inputDir := constants.TemplatesDir
	if len(cfg.Template.Path) > 0 {
		inputDir = filepath.Join(inputDir, cfg.Template.Path)
	}

	stat, err := os.Stat(inputDir)
	if err != nil {
		log.Fatal(fmt.Errorf("build: read template dir: %v", err))
	}

	if !stat.IsDir() {
		log.Fatal("build: read template dir: isn't dir")
	}

	err = copyDir(inputDir, cfg.OutputDir)
	if err != nil {
		log.Fatal(fmt.Errorf("build: copy template: %v", err))
	}

	log.Println("build: copy template: finish")
}

func renderTemplate(info *models.UserInfo) {
	log.Println("build: render template: start")

	cfg := config.GetConfig()

	if len(cfg.Template.Files) == 0 {
		log.Fatal("build: render template: template file name not specified")
	}

	for _, filename := range cfg.Template.Files {
		renderTemplateFile(filename, info)
	}

	log.Println("build: render template: finish")
}

func renderTemplateFile(filename string, info *models.UserInfo) {
	cfg := config.GetConfig()

	log.Printf("build: render template: render %v", filepath.Join(cfg.OutputDir, filename))

	f, err := os.Create(filepath.Join(cfg.OutputDir, filename))
	if err != nil {
		log.Fatal(fmt.Errorf("build: create template file: %v", err))
	}
	defer f.Close()

	t, err := template.ParseFiles(filepath.Join(constants.TemplatesDir, cfg.Template.Path, filename))
	if err != nil {
		log.Fatal(fmt.Errorf("build: parsing template: %v", err))
	}

	err = t.Execute(f, info)
	if err != nil {
		log.Fatal(fmt.Errorf("build: render template: %v", err))
	}
}

func copyDir(src, dst string) error {
	err := filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		localPath, err := filepath.Rel(src, srcPath)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, localPath)

		// Copy dir
		if info.IsDir() {
			err = os.MkdirAll(dstPath, info.Mode())
			return err
		}

		// Copy file
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

		err = dstFile.Chmod(info.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(dstFile, srcFile)
		return err
	})

	return err
}
