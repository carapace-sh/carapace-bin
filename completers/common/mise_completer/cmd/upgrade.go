package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrades outdated tools",
	Aliases: []string{"up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Show what would be upgraded without actually upgrading")
	upgradeCmd.Flags().BoolP("force", "f", false, "Force upgrade")
	upgradeCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	upgradeCmd.Flags().Bool("raw", false, "Connect backend install command directly to terminal")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
