package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var chooseBufferCmd = &cobra.Command{
	Use:   "choose-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseBufferCmd).Standalone()

	chooseBufferCmd.Flags().StringS("F", "F", "", "format")
	chooseBufferCmd.Flags().StringS("K", "K", "", "key-format")
	chooseBufferCmd.Flags().StringS("O", "O", "", "sort-order")
	chooseBufferCmd.Flags().StringS("f", "f", "", "filter")
	chooseBufferCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(chooseBufferCmd)
}
