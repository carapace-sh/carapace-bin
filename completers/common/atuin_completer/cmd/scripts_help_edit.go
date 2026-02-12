package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scripts_help_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_help_editCmd).Standalone()

	scripts_helpCmd.AddCommand(scripts_help_editCmd)
}
