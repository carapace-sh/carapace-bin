package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var renameWindowCmd = &cobra.Command{
	Use:   "rename-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameWindowCmd).Standalone()

	renameWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(renameWindowCmd)
}
