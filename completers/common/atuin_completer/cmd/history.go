package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Manipulate shell history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(historyCmd)
}
