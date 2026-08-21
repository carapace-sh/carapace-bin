package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade <package-name?>",
	Aliases: []string{"up"},
	Short:   "perform a full system upgrade, or update the specified packages along with any resulting dependencies",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("bypass-update-repo", "b", false, "do not update the repositories first")
	upgradeCmd.Flags().StringP("component", "c", "", "only upgrade packages from the given component")
	upgradeCmd.Flags().BoolP("dry-run", "n", false, "only show what would happen, do not actually perform changes")
	upgradeCmd.Flags().StringP("exclude", "x", "", "ignore packages and components that match the specified basename")
	upgradeCmd.Flags().String("exclude-from", "", "just like --exclude, except the list is specified in the given filename")
	upgradeCmd.Flags().BoolP("fetch-only", "f", false, "only download the packages, but do not apply any upgrade operations")
	upgradeCmd.Flags().Bool("ignore-build-no", false, "ignore build number errors")
	upgradeCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	upgradeCmd.Flags().Bool("ignore-file-conflicts", false, "allow completing the update even if file conflicts would occur")
	upgradeCmd.Flags().Bool("ignore-package-conflicts", false, "allow completing the upgrade even if package conflicts would occur")
	upgradeCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	upgradeCmd.Flags().StringP("repository", "r", "", "only upgrade packages from the given repository")
	upgradeCmd.Flags().Bool("security-only", false, "only apply updates that have been marked as security updates")

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		eopkg.ActionPackages().FilterArgs(),
	)
}
