package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Install upgrades available from repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "system maintenance",
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("available", "a", false, "Reset all packages to versions available from current repositories")
	upgradeCmd.Flags().Bool("ignore", false, "Upgrade all other packages than the ones listed")
	upgradeCmd.Flags().BoolP("latest", "l", false, "Always choose the latest package by version")
	upgradeCmd.Flags().Bool("no-self-upgrade", false, "Do not do an early upgrade of the 'apk-tools' package")
	upgradeCmd.Flags().Bool("prune", false, "Prune the WORLD by removing packages which are no longer available")
	upgradeCmd.Flags().Bool("self-upgrade-only", false, "Only perform a self-upgrade of the 'apk-tools' package")
	common.AddGlobalFlags(upgradeCmd)
	common.AddCommitFlags(upgradeCmd)
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		action.ActionPackages(upgradeCmd, true).FilterArgs(),
	)
}
