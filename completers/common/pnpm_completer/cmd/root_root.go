package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootRootCmd = &cobra.Command{
	Use:   "root",
	Short: "Print the effective modules directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootRootCmd).Standalone()

	rootRootCmd.Flags().BoolP("help", "h", false, "Output usage information")

	rootCmd.AddCommand(rootRootCmd)
}
