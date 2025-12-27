package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pamac"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "Upgrade your system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().BoolP("aur", "a", false, "also upgrade packages installed from AUR")
	updateCmd.Flags().String("builddir", "", "build directory (use with --aur)")
	updateCmd.Flags().Bool("devel", false, "also upgrade development packages (use with --aur)")
	updateCmd.Flags().Bool("disable-downgrade", false, "disable package downgrades")
	updateCmd.Flags().BoolP("download-only", "w", false, "download all packages but do not install/upgrade")
	updateCmd.Flags().Bool("enable-downgrade", false, "enable package downgrades")
	updateCmd.Flags().Bool("force-refresh", false, "force the refresh of the databases")
	updateCmd.Flags().String("ignore", "", "ignore a package upgrade")
	updateCmd.Flags().Bool("no-aur", false, "do not upgrade packages installed from AUR")
	updateCmd.Flags().Bool("no-confirm", false, "bypass any and all confirmation messages")
	updateCmd.Flags().Bool("no-devel", false, "do not upgrade development packages")
	updateCmd.Flags().String("overwrite", "", "overwrite conflicting files")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"builddir": carapace.ActionDirectories(),
		"ignore":   pamac.ActionInstalledPackages(false).UniqueList(","),
	})
}
