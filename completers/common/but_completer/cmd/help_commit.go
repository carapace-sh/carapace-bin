package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_commitCmd).Standalone()

	helpCmd.AddCommand(help_commitCmd)
}
