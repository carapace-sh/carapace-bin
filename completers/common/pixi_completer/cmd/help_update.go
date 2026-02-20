package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "The `update` command checks if there are newer versions of the dependencies and updates the `pixi.lock` file and environments accordingly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_updateCmd).Standalone()

	helpCmd.AddCommand(help_updateCmd)
}
