package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a revision in two",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	splitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to split. This is the default if no paths are provided")
	splitCmd.Flags().StringP("revision", "r", "", "The revision to split")
	rootCmd.AddCommand(splitCmd)
}
