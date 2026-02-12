package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_initStoreCmd = &cobra.Command{
	Use:   "init-store",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_initStoreCmd).Standalone()

	history_helpCmd.AddCommand(history_help_initStoreCmd)
}
