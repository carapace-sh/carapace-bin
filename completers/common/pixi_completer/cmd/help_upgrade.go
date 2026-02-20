package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Checks if there are newer versions of the dependencies and upgrades them in the lockfile and manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upgradeCmd).Standalone()

	helpCmd.AddCommand(help_upgradeCmd)
}
