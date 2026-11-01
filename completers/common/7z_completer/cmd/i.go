package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iCmd = &cobra.Command{
	Use:   "i",
	Short: "Show information about supported formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iCmd).Standalone()

	rootCmd.AddCommand(iCmd)
}
