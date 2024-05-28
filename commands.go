package main

import (
	"fmt"
	"path"
)

type CreateCmd struct {
	Name     string `arg:"-n,--name" help:"Name of the project"`
	Template string `arg:"-t,--template" help:"Template name to use"`
}

func (c *CreateCmd) Run(ctx *CommandContext) {
	createProject(ctx, c.Name, c.Template, ctx.conf.ProjectsDirectory)
}

type DropCmd struct {
	Name     string `arg:"-n,--name" help:"Name of the project to drop a template into"`
	Template string `arg:"-t,--template" help:"Template name to drop"`
	Hard     bool   `arg:"-h,--hard" help:"Overwrite everything if the template conflicts"`
}

type PadCmd struct {
	Template string `arg:"-t,--template" help:"Template name"`
}

func (c *PadCmd) Run(ctx *CommandContext) {
  
	createProject(ctx, "something", c.Template, ctx.conf.ScratchDirectory)
}

func (c *DropCmd) Run(ctx *CommandContext) {
  if c.Hard {
    fmt.Println("Warning! You are about to do a hard airdrop. This will overwrite any existing files. Enter 'y' if you are ok with these consequences.")

    var response string
    fmt.Scanln(&response)
    if response != "y" {
      fmt.Println("Aborting.")
      return
    } else {
      fmt.Println("Continuing with hard airdrop.")
    }
  }

  // Todo: implement the rest of the drop command 

}


// helper function used in create and pad commands
func createProject(ctx *CommandContext, name string, template string, directory string) {
	if name == "" {
		fmt.Println("Project name is required with -n. Aborting.")
		return
	}

	if template == "" {
		fmt.Println("Template name is required with -t. Aborting.")
		return
	}

	if ctx.fs.Exists(path.Join(directory, name)) {
		fmt.Println("Project with that name already exists. Aborting.")
		return
	}

	templatePath := path.Join(ctx.conf.TemplatesDirectory, template)
	if !ctx.fs.Exists(templatePath) {
		fmt.Println("Template not found. Aborting.")
		return
	}

	ctx.fs.Copy(templatePath, path.Join(directory, name))
	fmt.Println("Created project", name, "in", path.Join(directory, name))
}

type TemplatesCmd struct{}
type ProjectsCmd struct{}
type PadsCmd struct{}

type Elevate struct {
	Name string `arg:"-n,--name" default:"__none__" help:"Name of the project"`
}

type Archive struct {
	Name string `arg:"-n,--name" help:"Name of the project or scratch to archive"`
	All  bool   `arg:"-a,--all" help:"Archive all projects"`
}

type Unarchive struct {
	Name string `arg:"-n,--name" help:"Name of the project or scratch to unarchive"`
}

type Command interface {
	Run(ctx *CommandContext)
}
