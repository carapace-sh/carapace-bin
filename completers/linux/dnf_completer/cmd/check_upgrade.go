package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var check_upgradeCmd = &cobra.Command{
	Use:   "check-upgrade [options] [<package-spec>...]",
	Short: "check for available package upgrades",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(check_upgradeCmd).Standalone()

	check_upgradeCmd.Flags().Bool("changelogs", false, "print the package changelogs")
	check_upgradeCmd.Flags().Bool("json", false, "request JSON output format")
	check_upgradeCmd.Flags().Bool("minimal", false, "reports the lowest versions of packages that fix advisories")

	rootCmd.AddCommand(check_upgradeCmd)
}
