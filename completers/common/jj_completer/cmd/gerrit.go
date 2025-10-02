package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gerritCmd = &cobra.Command{
	Use:   "gerrit",
	Short: "Interact with Gerrit Code Review",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gerritCmd).Standalone()

	gerritCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(gerritCmd)
}
