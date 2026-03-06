package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade [options] [<package-spec>|@<group-spec>|@<environment-spec>...]",
	Aliases: []string{"up"},
	Short:   "upgrade packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("minimal", false, "upgrade packages to lowest versions fixing advisories")
	upgradeCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	upgradeCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	upgradeCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	upgradeCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	upgradeCmd.Flags().String("destdir", "", "set directory for downloading packages")
	upgradeCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(upgradeCmd),
	)
}
