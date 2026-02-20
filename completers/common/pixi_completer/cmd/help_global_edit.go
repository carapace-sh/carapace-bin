package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the global manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_editCmd).Standalone()

	help_globalCmd.AddCommand(help_global_editCmd)
}
