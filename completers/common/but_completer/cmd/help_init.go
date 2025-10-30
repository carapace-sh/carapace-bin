package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a GitButler project from a git repository in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_initCmd).Standalone()

	helpCmd.AddCommand(help_initCmd)
}
