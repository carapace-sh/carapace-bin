package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install packages",
	Aliases: []string{"it"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	installCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")
	installCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	installCmd.Flags().BoolP("dry-run", "n", false, "Do not perform any action, just show what would be done")
	installCmd.Flags().Bool("ignore-build-no", false, "Do not take build no into account")
	installCmd.Flags().Bool("reinstall", false, "Reinstall already installed packages")
	installCmd.Flags().Bool("ignore-check", false, "Skip distribution/version check")
	installCmd.Flags().Bool("ignore-file-conflicts", false, "Ignore file conflicts")
	installCmd.Flags().Bool("ignore-package-conflicts", false, "Ignore package conflicts")
	installCmd.Flags().StringP("component", "c", "", "Install component's packages")
	installCmd.Flags().StringP("repository", "r", "", "Name of the to be searched repository")
	installCmd.Flags().Bool("fetch-only", false, "Fetch upgrades but do not install")
	installCmd.Flags().String("exclude", "", "Packages to exclude from install")
	installCmd.Flags().String("exclude-from", "", "File containing packages to exclude")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".eopkg"),
	)
}
