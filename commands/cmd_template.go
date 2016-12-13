package commands

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ConfigureTemplateCommand setup "template" command
func ConfigureTemplateCommand(app *kingpin.Application) {
	template := app.Command("template", "Template engine")
	template.Command("install", "Install CV site template from config file").Action(runInstallCommand)
}

func runInstallCommand(*kingpin.ParseContext) error {
	log.Println("template: install template: start")

	cfg := GetConfig()

	if len(cfg.Template.RepoURL) > 0 {
		cloneRepo(cfg.Template)
	} else {
		log.Fatal("template: install template: repo URL not specefied")
	}

	log.Println("template: install template: finish")

	return nil
}

func cloneRepo(repo templateRepo) {
	log.Println("template: clone repo: start")

	r := git.NewMemoryRepository()

	err := r.Clone(&git.CloneOptions{URL: repo.RepoURL})
	if err != nil {
		log.Fatal(fmt.Errorf("template: clone repo: %v", err))
	}

	ref, err := r.Head()
	if err != nil {
		log.Fatal(fmt.Errorf("template: clone repo: %v", err))
	}

	commit, err := r.Commit(ref.Hash())
	if err != nil {
		log.Fatal(fmt.Errorf("template: clone repo: %v", err))
	}

	files, err := commit.Files()
	if err != nil {
		log.Fatal(fmt.Errorf("template: clone repo: %v", err))
	}

	err = files.ForEach(func(f *git.File) error {
		if len(repo.Path) > 0 && !strings.HasPrefix(f.Name, repo.Path) {
			return nil
		}

		abs := filepath.Join(templatesDir, f.Name)
		dir := filepath.Dir(abs)

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("template: create dir: %v", err)
		}

		file, err := os.Create(abs)
		if err != nil {
			return fmt.Errorf("template: create file: %v", err)
		}
		defer file.Close()

		r, err := f.Reader()
		if err != nil {
			return fmt.Errorf("template: read file: %v", err)
		}
		defer r.Close()

		if err := file.Chmod(f.Mode); err != nil {
			return fmt.Errorf("template: chmod: %v", err)
		}

		_, err = io.Copy(file, r)
		if err != nil {
			return fmt.Errorf("template: copy file: %v", err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("template: clone repo: finish")
}
