package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_absorbCmd = &cobra.Command{
	Use:   "absorb",
	Short: "Assign staged hunks into recent commits and replay descendants once",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_absorbCmd).Standalone()

	help_historyCmd.AddCommand(help_history_absorbCmd)
}
