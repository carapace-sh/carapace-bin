package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var update_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a new version of the GitButler CLI is available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(update_checkCmd).Standalone()

	update_checkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	updateCmd.AddCommand(update_checkCmd)
}
