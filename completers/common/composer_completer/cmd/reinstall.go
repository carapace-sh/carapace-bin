package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Uninstalls and reinstalls the given package names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().Bool("apcu-autoloader", false, "Use APCu to cache found/not-found classes.")
	reinstallCmd.Flags().String("apcu-autoloader-prefix", "", "Use a custom prefix for the APCu autoloader cache. Implicitly enables --apcu-autoloader")
	reinstallCmd.Flags().BoolP("classmap-authoritative", "a", false, "Autoload classes from the classmap only. Implicitly enables `--optimize-autoloader`.")
	reinstallCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	reinstallCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	reinstallCmd.Flags().Bool("no-autoloader", false, "Skips autoloader generation")
	reinstallCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	reinstallCmd.Flags().BoolP("optimize-autoloader", "o", false, "Optimize autoloader during autoloader dump")
	reinstallCmd.Flags().Bool("prefer-dist", false, "Forces installation from package dist (default behavior).")
	reinstallCmd.Flags().String("prefer-install", "", "Forces installation from package dist|source|auto (auto chooses source for dev versions, dist for the rest).")
	reinstallCmd.Flags().Bool("prefer-source", false, "Forces installation from package sources when possible, including VCS information.")
	reinstallCmd.Flags().StringArray("type", nil, "Filter packages to reinstall by type(s)")

	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).FlagCompletion(carapace.ActionMap{
		"ignore-platform-req": action.ActionPlatformPackages(reinstallCmd),
		"prefer-install":      carapace.ActionValues("auto", "dist", "source"),
		"type":                action.ActionInstalledPackageTypes(reinstallCmd),
	})

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(reinstallCmd).FilterArgs(),
	)
}
