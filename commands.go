package main

type CreateCmd struct {
	Name     string `arg:"positional" help:"Name of the project"`
	Template string `arg:"positional" help:"Template name to use"`
}

func (c *CreateCmd) Run(ctx *CommandContext) error {
  return nil
}

type PadCmd struct {
	Template string `arg:"positional" help:"Template name"`
}

func (c *PadCmd) Run(ctx *CommandContext) error {
  return nil
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
