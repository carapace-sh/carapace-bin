package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "List remotes, or add one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_remoteCmd).Standalone()

	helpCmd.AddCommand(help_remoteCmd)
}
