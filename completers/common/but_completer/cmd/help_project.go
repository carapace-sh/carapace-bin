package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "List and manipulate projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_projectCmd).Standalone()

	helpCmd.AddCommand(help_projectCmd)
}
