package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Removes a package from the require or require-dev",
	Aliases: []string{"rm", "uninstall"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("apcu-autoloader", false, "Use APCu to cache found/not-found classes.")
	removeCmd.Flags().String("apcu-autoloader-prefix", "", "Use a custom prefix for the APCu autoloader cache. Implicitly enables --apcu-autoloader")
	removeCmd.Flags().String("audit-format", "", "Audit output format. Must be \"table\", \"plain\", \"json\", or \"summary\".")
	removeCmd.Flags().BoolP("classmap-authoritative", "a", false, "Autoload classes from the classmap only. Implicitly enables `--optimize-autoloader`.")
	removeCmd.Flags().Bool("dev", false, "Removes a package from the require-dev section.")
	removeCmd.Flags().Bool("dry-run", false, "Outputs the operations but will not execute anything (implicitly enables --verbose).")
	removeCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	removeCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	removeCmd.Flags().BoolP("minimal-changes", "m", false, "During an update with -w/-W, only perform absolutely necessary changes to transitive dependencies (can also be set via the COMPOSER_MINIMAL_CHANGES=1 env var).")
	removeCmd.Flags().Bool("no-audit", false, "Skip the audit step after updating the composer.lock file (can also be set via the COMPOSER_NO_AUDIT=1 env var).")
	removeCmd.Flags().Bool("no-blocking", false, "Disables all policy blocking during this command (can also be set via the COMPOSER_NO_BLOCKING=1 env var).")
	removeCmd.Flags().Bool("no-install", false, "Skip the install step after updating the composer.lock file.")
	removeCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	removeCmd.Flags().Bool("no-security-blocking", false, "DEPRECATED: use --no-blocking instead. Allows installing packages with security advisories or that are abandoned (can also be set via the COMPOSER_NO_SECURITY_BLOCKING=1 env var).")
	removeCmd.Flags().Bool("no-update", false, "Disables the automatic update of the dependencies (implies --no-install).")
	removeCmd.Flags().Bool("no-update-with-dependencies", false, "Does not allow inherited dependencies to be updated with explicit dependencies.")
	removeCmd.Flags().BoolP("optimize-autoloader", "o", false, "Optimize autoloader during autoloader dump")
	removeCmd.Flags().Bool("unused", false, "Remove all packages which are locked but not required by any other package.")
	removeCmd.Flags().Bool("update-no-dev", false, "Run the dependency update with the --no-dev option.")
	removeCmd.Flags().BoolP("update-with-all-dependencies", "W", false, "Allows all inherited dependencies to be updated, including those that are root requirements (can also be set via the COMPOSER_WITH_ALL_DEPENDENCIES=1 env var).")
	removeCmd.Flags().BoolP("update-with-dependencies", "w", false, "Allows inherited dependencies to be updated with explicit dependencies (can also be set via the COMPOSER_WITH_DEPENDENCIES=1 env var). (Deprecated, is now default behavior)")
	removeCmd.Flags().Bool("with-all-dependencies", false, "Alias for --update-with-all-dependencies")

	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"audit-format":        carapace.ActionValues("json", "plain", "summary", "table"),
		"ignore-platform-req": action.ActionPlatformPackages(removeCmd),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		action.ActionRootRequirements(removeCmd).FilterArgs(),
	)
}
