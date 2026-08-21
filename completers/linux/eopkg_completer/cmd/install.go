package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install <name>",
	Aliases: []string{"it"},
	Short:   "install a named package or local .eopkg directly onto the system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringP("component", "c", "", "install an entire component by name, instead of just a package")
	installCmd.Flags().BoolP("dry-run", "n", false, "only show what would happen, do not actually perform changes")
	installCmd.Flags().StringP("exclude", "x", "", "ignore packages and components that match the specified basename")
	installCmd.Flags().String("exclude-from", "", "just like --exclude, except the list is specified in the given filename")
	installCmd.Flags().BoolP("fetch-only", "f", false, "download the required packages but don't actually install them")
	installCmd.Flags().Bool("ignore-build-no", false, "ignore build number errors")
	installCmd.Flags().Bool("ignore-check", false, "do not check if this package is intended for use with the current distribution")
	installCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	installCmd.Flags().Bool("ignore-dependency", false, "do not attempt the installation/validation of dependencies")
	installCmd.Flags().Bool("ignore-file-conflicts", false, "allow the package to install even if it conflicts with another package's files")
	installCmd.Flags().Bool("ignore-package-conflicts", false, "forcibly install a package even though it is marked as conflicting")
	installCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	installCmd.Flags().Bool("reinstall", false, "reinstall an already installed package")
	installCmd.Flags().StringP("repository", "r", "", "specify which repository to pull the component from")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles("eopkg").UnlessF(condition.CompletingPath),
			eopkg.ActionPackageSearch().UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
