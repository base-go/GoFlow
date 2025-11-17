package commands

import (
	"fmt"

	"github.com/base-go/mamba"
)

const CLIVersion = "0.2.0"

// VersionCommand creates the 'goflow version' command
func VersionCommand() *mamba.Command {
	cmd := &mamba.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show GoFlow CLI version",
		Long:    "Displays the current version of the GoFlow CLI.",
		Run: func(c *mamba.Command, args []string) {
			if err := runVersionCommand(c, args); err != nil {
				c.PrintError(fmt.Sprintf("Error: %v", err))
			}
		},
	}

	return cmd
}

func runVersionCommand(cmd *mamba.Command, args []string) error {
	cmd.PrintInfo(fmt.Sprintf("GoFlow CLI version %s", CLIVersion))
	return nil
}
