package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Mark a commit or branch for auto-assign or auto-commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_markCmd).Standalone()

	helpCmd.AddCommand(help_markCmd)
}
