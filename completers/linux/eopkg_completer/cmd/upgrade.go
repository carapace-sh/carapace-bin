package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade packages",
	Aliases: []string{"up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	upgradeCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")
	upgradeCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Do not perform any action, just show what would be done")
	upgradeCmd.Flags().Bool("ignore-build-no", false, "Do not take build no into account")
	upgradeCmd.Flags().Bool("security-only", false, "Security related upgrades only")
	upgradeCmd.Flags().BoolP("bypass-update-repo", "b", false, "Do not update repositories")
	upgradeCmd.Flags().Bool("ignore-file-conflicts", false, "Ignore file conflicts")
	upgradeCmd.Flags().Bool("ignore-package-conflicts", false, "Ignore package conflicts")
	upgradeCmd.Flags().StringP("component", "c", "", "Upgrade component's packages")
	upgradeCmd.Flags().StringP("repository", "r", "", "Name of the repository")
	upgradeCmd.Flags().Bool("fetch-only", false, "Fetch upgrades but do not install")
	upgradeCmd.Flags().String("exclude", "", "Packages to exclude from upgrade")
	upgradeCmd.Flags().String("exclude-from", "", "File containing packages to exclude")

	rootCmd.AddCommand(upgradeCmd)
}
