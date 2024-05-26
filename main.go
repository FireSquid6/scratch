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



func main() {
	fmt.Println("Hello, world!")

  ctx := GetContext()

  switch {
  case args.Create != nil:
    args.Create.Run(&ctx)
  case args.Pad != nil:
    args.Pad.Run(&ctx)
  default:
    fmt.Println("No comamnd specified. Try --help for more information.")
  }
}
