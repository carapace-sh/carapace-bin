package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_peersCmd = &cobra.Command{
	Use:   "peers",
	Short: "Checks for unmet or missing peer dependency issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_peersCmd).Standalone()

	helpCmd.AddCommand(help_peersCmd)
}
