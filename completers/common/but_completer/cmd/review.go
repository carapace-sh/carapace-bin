package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:     "review",
	Short:   "Commands for creating and publishing code reviews to a forge",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(reviewCmd).Standalone()

	reviewCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(reviewCmd)
}
