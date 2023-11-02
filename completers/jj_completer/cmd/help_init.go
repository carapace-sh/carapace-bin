package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new repo in the given directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_initCmd).Standalone()

	helpCmd.AddCommand(help_initCmd)
}
