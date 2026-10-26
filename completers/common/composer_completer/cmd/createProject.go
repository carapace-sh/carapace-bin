package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates new project from a package into given directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createProjectCmd).Standalone()

	createProjectCmd.Flags().Bool("add-repository", false, "Add the custom repository in the composer.json. If a lock file is present it will be deleted and an update will be run instead of install.")
	createProjectCmd.Flags().Bool("ask", false, "Whether to ask for project directory.")
	createProjectCmd.Flags().String("audit-format", "", "Audit output format. Must be \"table\", \"plain\", \"json\" or \"summary\".")
	createProjectCmd.Flags().Bool("dev", false, "Enables installation of require-dev packages (enabled by default, only present for BC).")
	createProjectCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages).")
	createProjectCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages).")
	createProjectCmd.Flags().Bool("keep-vcs", false, "Whether to prevent deleting the vcs folder.")
	createProjectCmd.Flags().Bool("no-audit", false, "Whether to skip auditing of the installed package dependencies (can also be set via the COMPOSER_NO_AUDIT=1 env var).")
	createProjectCmd.Flags().Bool("no-blocking", false, "Disables all policy blocking during this command (can also be set via the COMPOSER_NO_BLOCKING=1 env var).")
	createProjectCmd.Flags().Bool("no-custom-installers", false, "DEPRECATED: Use no-plugins instead.")
	createProjectCmd.Flags().Bool("no-dev", false, "Disables installation of require-dev packages.")
	createProjectCmd.Flags().Bool("no-install", false, "Whether to skip installation of the package dependencies.")
	createProjectCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	createProjectCmd.Flags().Bool("no-secure-http", false, "Disable the secure-http config option temporarily while installing the root package. Use at your own risk. Using this flag is a bad idea.")
	createProjectCmd.Flags().Bool("no-security-blocking", false, "DEPRECATED: use --no-blocking instead. Allows installing packages with security advisories or that are abandoned (can also be set via the COMPOSER_NO_SECURITY_BLOCKING=1 env var).")
	createProjectCmd.Flags().Bool("prefer-dist", false, "Forces installation from package dist (default behavior).")
	createProjectCmd.Flags().String("prefer-install", "", "Forces installation from package dist|source|auto (auto chooses source for dev versions, dist for the rest).")
	createProjectCmd.Flags().Bool("prefer-source", false, "Forces installation from package sources when possible, including VCS information.")
	createProjectCmd.Flags().Bool("remove-vcs", false, "Whether to force deletion of the vcs folder without prompting.")
	createProjectCmd.Flags().StringArray("repository", nil, "Add custom repositories to look the package up, either by URL or using JSON arrays")
	createProjectCmd.Flags().String("repository-url", "", "DEPRECATED: Use --repository instead.")
	createProjectCmd.Flags().StringArray("require", nil, "Require additional package(s) to composer.json after installing the project. If a lock file is present it will be deleted and an update will be run instead of install.")
	createProjectCmd.Flags().StringP("stability", "s", "", "Minimum-stability allowed (unless a version is specified).")

	rootCmd.AddCommand(createProjectCmd)

	carapace.Gen(createProjectCmd).FlagCompletion(carapace.ActionMap{
		"audit-format":        carapace.ActionValues("json", "plain", "summary", "table"),
		"ignore-platform-req": action.ActionPlatformPackages(createProjectCmd),
		"prefer-install":      carapace.ActionValues("auto", "dist", "source"),
		"stability":           carapace.ActionValues("RC", "alpha", "beta", "dev", "stable"),
	})

	carapace.Gen(createProjectCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
