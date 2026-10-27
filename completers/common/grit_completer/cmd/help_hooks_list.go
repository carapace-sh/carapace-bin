package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hooks_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List hook scripts in the resolved hooks directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hooks_listCmd).Standalone()

	help_hooksCmd.AddCommand(help_hooks_listCmd)
}
