package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_initStoreCmd = &cobra.Command{
	Use:   "init-store",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_initStoreCmd).Standalone()

	help_historyCmd.AddCommand(help_history_initStoreCmd)
}
