package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Installs the project dependencies from the composer.lock file if present, or falls back on the composer.json",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("apcu-autoloader", false, "Use APCu to cache found/not-found classes.")
	installCmd.Flags().String("apcu-autoloader-prefix", "", "Use a custom prefix for the APCu autoloader cache. Implicitly enables --apcu-autoloader")
	installCmd.Flags().Bool("audit", false, "Run an audit after installation is complete.")
	installCmd.Flags().String("audit-format", "", "Audit output format. Must be \"table\", \"plain\", \"json\", or \"summary\".")
	installCmd.Flags().BoolP("classmap-authoritative", "a", false, "Autoload classes from the classmap only. Implicitly enables `--optimize-autoloader`.")
	installCmd.Flags().Bool("dev", false, "DEPRECATED: Enables installation of require-dev packages (enabled by default, only present for BC).")
	installCmd.Flags().Bool("download-only", false, "Download only, do not install packages.")
	installCmd.Flags().Bool("dry-run", false, "Outputs the operations but will not execute anything (implicitly enables --verbose).")
	installCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	installCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	installCmd.Flags().Bool("no-autoloader", false, "Skips autoloader generation")
	installCmd.Flags().Bool("no-blocking", false, "Disables all policy blocking during this command (can also be set via the COMPOSER_NO_BLOCKING=1 env var).")
	installCmd.Flags().Bool("no-dev", false, "Disables installation of require-dev packages.")
	installCmd.Flags().Bool("no-install", false, "Do not use, only defined here to catch misuse of the install command.")
	installCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	installCmd.Flags().Bool("no-security-blocking", false, "DEPRECATED: use --no-blocking instead. Allows installing packages with security advisories or that are abandoned (can also be set via the COMPOSER_NO_SECURITY_BLOCKING=1 env var).")
	installCmd.Flags().Bool("no-suggest", false, "DEPRECATED: This flag does not exist anymore.")
	installCmd.Flags().BoolP("optimize-autoloader", "o", false, "Optimize autoloader during autoloader dump")
	installCmd.Flags().Bool("prefer-dist", false, "Forces installation from package dist (default behavior).")
	installCmd.Flags().String("prefer-install", "", "Forces installation from package dist|source|auto (auto chooses source for dev versions, dist for the rest).")
	installCmd.Flags().Bool("prefer-source", false, "Forces installation from package sources when possible, including VCS information.")
	installCmd.Flags().Bool("strict-psr-autoloader", false, "Return a failed status code (6) if PSR-4 or PSR-0 mapping errors are present. Requires --optimize-autoloader to work.")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"audit-format":        carapace.ActionValues("json", "plain", "summary", "table"),
		"ignore-platform-req": action.ActionPlatformPackages(installCmd),
		"prefer-install":      carapace.ActionValues("auto", "dist", "source"),
	})
}
