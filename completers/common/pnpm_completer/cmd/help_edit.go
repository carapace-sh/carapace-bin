package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Not implemented in pnpm. Use the npm CLI directly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_editCmd).Standalone()

	helpCmd.AddCommand(help_editCmd)
}
