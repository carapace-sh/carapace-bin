package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ShowOptionCmd = &cobra.Command{
	Use:   "show-option",
	Short: "Show Option",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ShowOptionCmd).Standalone()
	rootCmd.AddCommand(ShowOptionCmd)
}
