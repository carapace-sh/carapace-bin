package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"U"},
	Short:   "Upgrade or add packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("install", "i", false, "Install package as well as building")

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
