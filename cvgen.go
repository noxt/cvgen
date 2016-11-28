package main

import (
	"github.com/noxt/cvgen/commands"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	kingpin.Version("0.0.1").Author("Dmitry Ivanenko")
	app := kingpin.New("cvgen", "CV generator from YAML files.")
	commands.ConfigureInitCommand(app)
	commands.ConfigureTemplateCommand(app)
	commands.ConfigureBuildCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
