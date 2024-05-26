package main

import (
	"fmt"
)

var args struct {
	Silent bool `default:"false" help:"Disable all output"`

	Create    *CreateCmd   `arg:"subcommand:create" help:"Create a new project"`
	Pad       *PadCmd      `arg:"subcommand:pad" help:"Create a scratchpad"`
	PadList   *PadsCmd     `arg:"subcommand:pads" help:"List scratchpads"`
	Template  *TemplateCmd `arg:"subcommand:template" help:"Create a new template"`
	Projects  *ProjectsCmd `arg:"subcommand:projects" help:"List projects"`
	Elevate   *Elevate     `arg:"subcommand:elevate" help:"Elevate a project to a template"`
	Archive   *Archive     `arg:"subcommand:archive" help:"Archive a project"`
	Unarchive *Unarchive   `arg:"subcommand:unarchive" help:"Unarchive a project"`
}

type CreateCmd struct {
	Name     string `arg:"positional" help:"Name of the project"`
	Template string `arg:"positional" help:"Template name to use"`
}

type PadCmd struct {
	Template string `arg:"positional" help:"Template name"`
}

type TemplateCmd struct {
	Name string `arg:"positional" help:"Name of the template"`
}

type ProjectsCmd struct{}
type PadsCmd struct{}

type Elevate struct {
	Name string `arg:"-n,--name" default:"__none__" help:"Name of the project"`
}

type Archive struct {
	Name string `arg:"positional" help:"Name of the project or scratch to archive"`
}

type Unarchive struct {
	Name string `arg:"positional" help:"Name of the project or scratch to unarchive"`
}

type Command interface {
	Run(ctx *CommandContext) error
}

func main() {
	fmt.Println("Hello, world!")
}
