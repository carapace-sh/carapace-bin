package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hooks_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List hook scripts in the resolved hooks directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hooks_help_listCmd).Standalone()

	hooks_helpCmd.AddCommand(hooks_help_listCmd)
}
