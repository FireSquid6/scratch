package main

import (
	"fmt"
	"math/rand"
	"os"
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
	Name     string `arg:"-n,--name" help:"Name of the project to drop a template into. Will be current directory if not specified"`
	Template string `arg:"-t,--template" help:"Template name to drop"`
}

type PadCmd struct {
	Template string `arg:"-t,--template" help:"Template name"`
}

func (c *PadCmd) Run(ctx *CommandContext) {
	nouns := getNouns()
	adjectives := getAdjectives()

	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	createProject(ctx, adjective+"-"+noun, c.Template, ctx.conf.ScratchDirectory)
}

func (c *DropCmd) Run(ctx *CommandContext) {
  fmt.Println("Warning! You are about to do a hard airdrop. This will overwrite any existing files. Enter 'y' if you are ok with these consequences.")

  var response string
  fmt.Scanln(&response)
  if response != "y" {
    fmt.Println("Aborting.")
    return
  } else {
    fmt.Println("Continuing with hard airdrop.")
  }

	// Todo: implement the rest of the drop command
	templatePath := path.Join(ctx.conf.TemplatesDirectory, c.Template)
	if !ctx.fs.Exists(templatePath) {
		fmt.Println("Template not found. Aborting.")
		return
	}

	filepath := ""
	err := error(nil)
	if c.Name == "" {
		filepath, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
		}
	} else {
		filepath = path.Join(ctx.conf.ProjectsDirectory, c.Name)
	}

	ctx.fs.Copy(templatePath, filepath)
  fmt.Println("Dropped template", c.Template, "into", filepath)
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

func (c *TemplatesCmd) Run(ctx *CommandContext) {
	filepaths, err := ctx.fs.GetFilepathsInDirectory(ctx.conf.TemplatesDirectory)

	if err != nil {
		fmt.Println("Error reading templates directory:", err)
		return
	}

	for _, filepath := range filepaths {
		_, file := path.Split(filepath)

		fmt.Println(file)
	}
}

type ProjectsCmd struct{}
type PadsCmd struct{}

func (c *PadsCmd) Run(ctx *CommandContext) {
	filepaths, err := ctx.fs.GetFilepathsInDirectory(ctx.conf.ScratchDirectory)
	if err != nil {
		fmt.Println("Error reading scratch directory:", err)
		return
	}

	for _, filepath := range filepaths {
		_, file := path.Split(filepath)

		fmt.Println(file)
	}
}

func (c *ProjectsCmd) Run(ctx *CommandContext) {
	filepaths, err := ctx.fs.GetFilepathsInDirectory(ctx.conf.ProjectsDirectory)

	if err != nil {
		fmt.Println("Error reading projects directory:", err)
		return
	}

	for _, filepath := range filepaths {
		_, file := path.Split(filepath)

		fmt.Println(file)
	}
}

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
