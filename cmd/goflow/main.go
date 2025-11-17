package main

import (
	"fmt"
	"os"

	"github.com/base-go/GoFlow/cmd/goflow/commands"
	"github.com/base-go/mamba"
)

func main() {
	rootCmd := &mamba.Command{
		Use:   "goflow",
		Short: "GoFlow - Flutter-inspired GUI framework for Go",
		Long: `GoFlow is a modern GUI framework for Go that brings Flutter's
developer experience to desktop and web applications.

Features:
  • Reactive UI with Signals
  • Hot reload for rapid development
  • Comprehensive widget library
  • Navigation & routing
  • Testing framework
  • Cross-platform support`,
		Version: commands.CLIVersion,
	}

	// Add modular commands
	rootCmd.AddCommand(commands.NewCommand())
	rootCmd.AddCommand(commands.RunCommand())
	rootCmd.AddCommand(commands.VersionCommand())

	// TODO: Migrate these commands to separate files
	// rootCmd.AddCommand(commands.BuildCommand())
	// rootCmd.AddCommand(commands.TestCommand())
	// rootCmd.AddCommand(commands.DoctorCommand())
	// rootCmd.AddCommand(commands.CleanCommand())
	// rootCmd.AddCommand(commands.AnalyzeCommand())
	// rootCmd.AddCommand(commands.FormatCommand())

	// Execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
