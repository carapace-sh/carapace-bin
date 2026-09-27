package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a package. Runs \"stop\", \"restart\" (if present), and \"start\" scripts, and associated pre- and post- scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_restartCmd).Standalone()

	helpCmd.AddCommand(help_restartCmd)
}
