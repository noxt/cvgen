package commands

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TemplateCommand struct {
}

func ConfigureTemplateCommand(app *kingpin.Application) {
	cmd := &TemplateCommand{}
	template := app.Command("template", "CV templates")
	template.Command("install", "Install templates from config file").Action(cmd.install)
}

func (cmd *TemplateCommand) install(c *kingpin.ParseContext) error {
	cfg := GetConfig()

	if len(cfg.Template.RepoURL) > 0 {
		err := cloneTemplateRepo(cfg.Template)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func cloneTemplateRepo(repo templateRepo) error {
	r := git.NewMemoryRepository()

	if err := r.Clone(&git.CloneOptions{URL: repo.RepoURL}); err != nil {
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
		if !strings.HasPrefix(f.Name, repo.Name) {
			return nil
		}

		abs := filepath.Join(TemplatesDir, f.Name)
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
		return err
	}

	return nil
}
