package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var requireCmd = &cobra.Command{
	Use:     "require",
	Short:   "Adds required packages to your composer.json and installs them",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(requireCmd).Standalone()

	requireCmd.Flags().Bool("apcu-autoloader", false, "Use APCu to cache found/not-found classes.")
	requireCmd.Flags().String("apcu-autoloader-prefix", "", "Use a custom prefix for the APCu autoloader cache. Implicitly enables --apcu-autoloader")
	requireCmd.Flags().String("audit-format", "", "Audit output format. Must be \"table\", \"plain\", \"json\", or \"summary\".")
	requireCmd.Flags().BoolP("classmap-authoritative", "a", false, "Autoload classes from the classmap only. Implicitly enables `--optimize-autoloader`.")
	requireCmd.Flags().Bool("dev", false, "Add requirement to require-dev.")
	requireCmd.Flags().Bool("dry-run", false, "Outputs the operations but will not execute anything (implicitly enables --verbose).")
	requireCmd.Flags().Bool("fixed", false, "Write fixed version to the composer.json.")
	requireCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	requireCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	requireCmd.Flags().BoolP("minimal-changes", "m", false, "During an update with -w/-W, only perform absolutely necessary changes to transitive dependencies (can also be set via the COMPOSER_MINIMAL_CHANGES=1 env var).")
	requireCmd.Flags().Bool("no-audit", false, "Skip the audit step after updating the composer.lock file (can also be set via the COMPOSER_NO_AUDIT=1 env var).")
	requireCmd.Flags().Bool("no-blocking", false, "Disables all policy blocking during this command (can also be set via the COMPOSER_NO_BLOCKING=1 env var).")
	requireCmd.Flags().Bool("no-install", false, "Skip the install step after updating the composer.lock file.")
	requireCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	requireCmd.Flags().Bool("no-security-blocking", false, "DEPRECATED: use --no-blocking instead. Allows installing packages with security advisories or that are abandoned (can also be set via the COMPOSER_NO_SECURITY_BLOCKING=1 env var).")
	requireCmd.Flags().Bool("no-suggest", false, "DEPRECATED: This flag does not exist anymore.")
	requireCmd.Flags().Bool("no-update", false, "Disables the automatic update of the dependencies (implies --no-install).")
	requireCmd.Flags().BoolP("optimize-autoloader", "o", false, "Optimize autoloader during autoloader dump")
	requireCmd.Flags().Bool("prefer-dist", false, "Forces installation from package dist (default behavior).")
	requireCmd.Flags().String("prefer-install", "", "Forces installation from package dist|source|auto (auto chooses source for dev versions, dist for the rest).")
	requireCmd.Flags().Bool("prefer-lowest", false, "Prefer lowest versions of dependencies (can also be set via the COMPOSER_PREFER_LOWEST=1 env var).")
	requireCmd.Flags().Bool("prefer-source", false, "Forces installation from package sources when possible, including VCS information.")
	requireCmd.Flags().Bool("prefer-stable", false, "Prefer stable versions of dependencies (can also be set via the COMPOSER_PREFER_STABLE=1 env var).")
	requireCmd.Flags().Bool("sort-packages", false, "Sorts packages when adding/updating a new dependency")
	requireCmd.Flags().Bool("update-no-dev", false, "Run the dependency update with the --no-dev option.")
	requireCmd.Flags().BoolP("update-with-all-dependencies", "W", false, "Allows all inherited dependencies to be updated, including those that are root requirements (can also be set via the COMPOSER_WITH_ALL_DEPENDENCIES=1 env var).")
	requireCmd.Flags().BoolP("update-with-dependencies", "w", false, "Allows inherited dependencies to be updated, except those that are root requirements (can also be set via the COMPOSER_WITH_DEPENDENCIES=1 env var).")
	requireCmd.Flags().Bool("with-all-dependencies", false, "Alias for --update-with-all-dependencies")
	requireCmd.Flags().Bool("with-dependencies", false, "Alias for --update-with-dependencies")

	rootCmd.AddCommand(requireCmd)

	carapace.Gen(requireCmd).FlagCompletion(carapace.ActionMap{
		"audit-format":        carapace.ActionValues("json", "plain", "summary", "table"),
		"ignore-platform-req": action.ActionPlatformPackages(requireCmd),
		"prefer-install":      carapace.ActionValues("auto", "dist", "source"),
	})
}
