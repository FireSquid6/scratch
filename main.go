package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
)

var args struct {
	Silent bool `default:"false" help:"Disable all output"`

	Create    *CreateCmd    `arg:"subcommand:create" help:"Create a new project"`
	Pad       *PadCmd       `arg:"subcommand:pad" help:"Create a scratchpad"`
	PadList   *PadsCmd      `arg:"subcommand:pads" help:"List scratchpads"`
	Projects  *ProjectsCmd  `arg:"subcommand:projects" help:"List projects"`
	Templates *TemplatesCmd `arg:"subcommand:templates" help:"List templates"`
	Elevate   *Elevate      `arg:"subcommand:elevate" help:"Elevate a project to a template"`
	Archive   *Archive      `arg:"subcommand:archive" help:"Archive a project"`
	Unarchive *Unarchive    `arg:"subcommand:unarchive" help:"Unarchive a project"`
	Airdrop   *DropCmd      `arg:"subcommand:airdrop" help:"Copy a file to the scratch directory"`
}

func main() {
	arg.MustParse(&args)
	ctx := GetContext()

	switch {
	case args.Create != nil:
		args.Create.Run(&ctx)
	case args.Pad != nil:
		args.Pad.Run(&ctx)
	case args.Templates != nil:
		args.Templates.Run(&ctx)
	case args.PadList != nil:
		args.PadList.Run(&ctx)
	case args.Projects != nil:
		args.Projects.Run(&ctx)
	case args.Airdrop != nil:
		args.Airdrop.Run(&ctx)
	default:
		fmt.Println("No comamnd specified. Try --help for more information.")
	}
}
