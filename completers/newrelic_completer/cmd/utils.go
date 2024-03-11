package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "Various utility methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utilsCmd).Standalone()
	rootCmd.AddCommand(utilsCmd)
}
