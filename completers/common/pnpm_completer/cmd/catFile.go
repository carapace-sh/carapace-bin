package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Prints the contents of a file based on the hash value stored in the index file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catFileCmd).Standalone()

	catFileCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(catFileCmd)
}
