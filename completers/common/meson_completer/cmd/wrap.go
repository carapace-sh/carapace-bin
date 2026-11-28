package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrapCmd = &cobra.Command{
	Use:   "wrap",
	Short: "Wrap tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrapCmd).Standalone()

	applyCommonFlagCompletion(wrapCmd)
	rootCmd.AddCommand(wrapCmd)
}
