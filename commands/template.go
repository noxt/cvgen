package commands

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"os"
	"path/filepath"
)

type TemplateCommand struct {
}

func (cmd *TemplateCommand) install(c *kingpin.ParseContext) error {
	cfg := GetConfig()

	r := git.NewMemoryRepository()

	err := r.Clone(&git.CloneOptions{
		URL:          cfg.Templates.Repo,
		SingleBranch: true,
	})
	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)

	commit, err := r.Commit(ref.Hash())
	CheckIfError(err)

	files, err := commit.Files()
	CheckIfError(err)

	err = files.ForEach(func(f *git.File) error {
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
	CheckIfError(err)

	return nil
}

func ConfigureTemplateCommand(app *kingpin.Application) {
	cmd := &TemplateCommand{}
	template := app.Command("template", "CV templates")
	template.Command("install", "Install templates from config file").Action(cmd.install)
}
