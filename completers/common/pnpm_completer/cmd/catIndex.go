package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catIndexCmd = &cobra.Command{
	Use:   "cat-index",
	Short: "Prints the index file of a specific package from the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catIndexCmd).Standalone()

	catIndexCmd.Flags().BoolP("help", "h", false, "Output usage information")

	rootCmd.AddCommand(catIndexCmd)
}
