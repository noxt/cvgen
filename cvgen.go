package main

import (
	"github.com/noxt/cvgen/commands"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func main() {
	app := kingpin.New("cvgen", "Generate CV sites with YAML files")
	app.Version("0.0.1").Author("Dmitry Ivanenko")
	commands.ConfigureInitCommand(app)
	commands.ConfigureTemplateCommand(app)
	commands.ConfigureBuildCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
