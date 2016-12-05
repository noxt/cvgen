package commands

import (
	"errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

const outputDir = "output"

// ConfigureBuildCommand setup "build" command
func ConfigureBuildCommand(app *kingpin.Application) {
	app.Command("build", "Build CV site from current directory").Action(runBuildCommand)
}

func runBuildCommand(*kingpin.ParseContext) error {
	var userInfo userInfo
	userInfo.load()
	userInfo.copyTemplate()
	userInfo.render()

	return nil
}

func (info *userInfo) load() {
	log.Println("Start parsing YAML files...")

	var parsingMap = map[string]interface{}{
		aboutMeFileName:       &info.AboutMe,
		educationFileName:     &info.Educations,
		organizationsFileName: &info.Organizations,
		projectsFileName:      &info.Projects,
		skillsFileName:        &info.Skills,
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

	log.Println("YAML files successfully parsed!")
}

func (info *userInfo) copyTemplate() {
	log.Println("Start copying template...")

	cfg := GetConfig()

	inputDir := templatesDir
	if len(cfg.Template.Path) > 0 {
		inputDir = filepath.Join(inputDir, cfg.Template.Path)
	}

	err := copyDir(inputDir, outputDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Template successfully copied!")
}

func (info *userInfo) render() {
	cfg := GetConfig()

	if len(cfg.Template.Files) == 0 {
		log.Fatal(errors.New("Template file name not specified"))
	}

	for _, filename := range cfg.Template.Files {
		log.Printf("Render template: %v\n", filepath.Join(outputDir, filename))

		f, err := os.Create(filepath.Join(outputDir, filename))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		t, err := template.ParseFiles(filepath.Join(templatesDir, cfg.Template.Path, filename))
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(f, info)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyDir(src, dst string) error {
	err := filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
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
