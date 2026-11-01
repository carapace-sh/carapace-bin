package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Updates your dependencies to the latest version according to composer.json, and updates the composer.lock file",
	Aliases: []string{"u", "upgrade"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("apcu-autoloader", false, "Use APCu to cache found/not-found classes.")
	updateCmd.Flags().String("apcu-autoloader-prefix", "", "Use a custom prefix for the APCu autoloader cache. Implicitly enables --apcu-autoloader")
	updateCmd.Flags().String("audit-format", "", "Audit output format. Must be \"table\", \"plain\", \"json\", or \"summary\".")
	updateCmd.Flags().String("bump-after-update", "", "Runs bump after performing the update.")
	updateCmd.Flags().BoolP("classmap-authoritative", "a", false, "Autoload classes from the classmap only. Implicitly enables `--optimize-autoloader`.")
	updateCmd.Flags().Bool("dev", false, "DEPRECATED: Enables installation of require-dev packages (enabled by default, only present for BC).")
	updateCmd.Flags().Bool("dry-run", false, "Outputs the operations but will not execute anything (implicitly enables --verbose).")
	updateCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	updateCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	updateCmd.Flags().BoolP("interactive", "i", false, "Interactive interface with autocompletion to select the packages to update.")
	updateCmd.Flags().Bool("lock", false, "Overwrites the lock file hash to suppress warning about the lock file being out of date without updating package versions. Package metadata like mirrors and URLs are updated if they changed.")
	updateCmd.Flags().BoolP("minimal-changes", "m", false, "Only perform absolutely necessary changes to dependencies. If packages cannot be kept at their currently locked version they are updated. For partial updates the allow-listed packages are always updated fully. (can also be set via the COMPOSER_MINIMAL_CHANGES=1 env var).")
	updateCmd.Flags().Bool("no-audit", false, "Skip the audit step after updating the composer.lock file (can also be set via the COMPOSER_NO_AUDIT=1 env var).")
	updateCmd.Flags().Bool("no-autoloader", false, "Skips autoloader generation")
	updateCmd.Flags().Bool("no-blocking", false, "Disables all policy blocking during this command (can also be set via the COMPOSER_NO_BLOCKING=1 env var).")
	updateCmd.Flags().Bool("no-dev", false, "Disables installation of require-dev packages.")
	updateCmd.Flags().Bool("no-install", false, "Skip the install step after updating the composer.lock file.")
	updateCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	updateCmd.Flags().Bool("no-security-blocking", false, "DEPRECATED: use --no-blocking instead. Allows installing packages with security advisories or that are abandoned (can also be set via the COMPOSER_NO_SECURITY_BLOCKING=1 env var).")
	updateCmd.Flags().Bool("no-suggest", false, "DEPRECATED: This flag does not exist anymore.")
	updateCmd.Flags().BoolP("optimize-autoloader", "o", false, "Optimize autoloader during autoloader dump.")
	updateCmd.Flags().Bool("patch-only", false, "Only allow patch version updates for currently installed dependencies.")
	updateCmd.Flags().Bool("prefer-dist", false, "Forces installation from package dist (default behavior).")
	updateCmd.Flags().String("prefer-install", "", "Forces installation from package dist|source|auto (auto chooses source for dev versions, dist for the rest).")
	updateCmd.Flags().Bool("prefer-lowest", false, "Prefer lowest versions of dependencies (can also be set via the COMPOSER_PREFER_LOWEST=1 env var).")
	updateCmd.Flags().Bool("prefer-source", false, "Forces installation from package sources when possible, including VCS information.")
	updateCmd.Flags().Bool("prefer-stable", false, "Prefer stable versions of dependencies (can also be set via the COMPOSER_PREFER_STABLE=1 env var).")
	updateCmd.Flags().Bool("root-reqs", false, "Restricts the update to your first degree dependencies.")
	updateCmd.Flags().Bool("strict-psr-autoloader", false, "Return a failed status code (6) if PSR-4 or PSR-0 mapping errors are present. Requires --optimize-autoloader to work.")
	updateCmd.Flags().StringArray("with", nil, "Temporary version constraint to add, e.g. foo/bar:1.0.0 or foo/bar=1.0.0")
	updateCmd.Flags().BoolP("with-all-dependencies", "W", false, "Update also dependencies of packages in the argument list, including those which are root requirements (can also be set via the COMPOSER_WITH_ALL_DEPENDENCIES=1 env var).")
	updateCmd.Flags().BoolP("with-dependencies", "w", false, "Update also dependencies of packages in the argument list, except those which are root requirements (can also be set via the COMPOSER_WITH_DEPENDENCIES=1 env var).")

	updateCmd.Flag("bump-after-update").NoOptDefVal = " "

	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"audit-format":        carapace.ActionValues("json", "plain", "summary", "table"),
		"bump-after-update":   carapace.ActionValues("all", "dev", "no-dev"),
		"ignore-platform-req": action.ActionPlatformPackages(updateCmd),
		"prefer-install":      carapace.ActionValues("auto", "dist", "source"),
	})

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(updateCmd).FilterArgs(),
	)
}
