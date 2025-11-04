package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Command for creating and publishing code reviews to a forge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reviewCmd).Standalone()

	reviewCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(reviewCmd)
}
