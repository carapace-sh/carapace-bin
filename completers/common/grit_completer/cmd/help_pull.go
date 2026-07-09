package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch from the remote and integrate it into the current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pullCmd).Standalone()

	helpCmd.AddCommand(help_pullCmd)
}
