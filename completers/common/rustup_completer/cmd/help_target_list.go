package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_target_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed and available targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_target_listCmd).Standalone()

	help_targetCmd.AddCommand(help_target_listCmd)
}
