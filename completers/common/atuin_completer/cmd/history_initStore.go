package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_initStoreCmd = &cobra.Command{
	Use:   "init-store",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_initStoreCmd).Standalone()

	history_initStoreCmd.Flags().BoolP("help", "h", false, "Print help")
	historyCmd.AddCommand(history_initStoreCmd)
}
