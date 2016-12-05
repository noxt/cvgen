package commands

import (
	"errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ConfigureTemplateCommand setup "template" command
func ConfigureTemplateCommand(app *kingpin.Application) {
	template := app.Command("template", "Template engine")
	template.Command("install", "Install CV site template from config file").Action(runInstallCommand)
}

func runInstallCommand(*kingpin.ParseContext) error {
	cfg := GetConfig()

	if len(cfg.Template.RepoURL) > 0 {
		err := cloneTemplateRepo(cfg.Template)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(errors.New("Template URL not specefied"))
	}

	return nil
}

func cloneTemplateRepo(repo templateRepo) error {
	r := git.NewMemoryRepository()

	err := r.Clone(&git.CloneOptions{URL: repo.RepoURL})
	if err != nil {
		return err
	}

	ref, err := r.Head()
	if err != nil {
		return err
	}

	commit, err := r.Commit(ref.Hash())
	if err != nil {
		return err
	}

	files, err := commit.Files()
	if err != nil {
		return err
	}

	err = files.ForEach(func(f *git.File) error {
		if len(repo.Path) > 0 && !strings.HasPrefix(f.Name, repo.Path) {
			return nil
		}

		abs := filepath.Join(templatesDir, f.Name)
		dir := filepath.Dir(abs)

		os.MkdirAll(dir, os.ModePerm)

		file, err := os.Create(abs)
		if err != nil {
			return err
		}
		defer file.Close()

		r, err := f.Reader()
		if err != nil {
			return err
		}
		defer r.Close()

		if err := file.Chmod(f.Mode); err != nil {
			return err
		}

		_, err = io.Copy(file, r)

		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	if len(repo.Path) > 0 {
		_, err := os.Stat(filepath.Join(templatesDir, repo.Path))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Template successfully cloned: %v\n", path.Join(repo.RepoURL, repo.Path))
		}
	} else {
		log.Printf("Template successfully cloned: %v\n", repo.RepoURL)
	}

	return nil
}
