package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:     "checkout",
	Short:   "Create a new, empty change and edit it in the working copy",
	Aliases: []string{"co"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkoutCmd).Standalone()

	checkoutCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	checkoutCmd.Flags().StringSliceP("message", "m", []string{}, "The change description to use")
	rootCmd.AddCommand(checkoutCmd)
}
