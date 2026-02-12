package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scripts_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_help_deleteCmd).Standalone()

	scripts_helpCmd.AddCommand(scripts_help_deleteCmd)
}
