package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Alter rustup settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_setCmd).Standalone()

	helpCmd.AddCommand(help_setCmd)
}
