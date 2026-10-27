package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove tracked paths from the working tree and index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rmCmd).Standalone()

	helpCmd.AddCommand(help_rmCmd)
}
