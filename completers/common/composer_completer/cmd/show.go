package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Shows information about packages",
	Aliases: []string{"info"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().Bool("all", false, "List all packages")
	showCmd.Flags().BoolP("available", "a", false, "List available packages only")
	showCmd.Flags().BoolP("direct", "D", false, "Shows only packages that are directly required by the root package")
	showCmd.Flags().StringP("format", "f", "", "Format of the output: text or json")
	showCmd.Flags().StringArray("ignore", nil, "Ignore specified package(s). Can contain wildcards (*). Use it with the --outdated option if you don't want to be informed about new versions of some packages.")
	showCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages). Use with the --outdated option")
	showCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages). Use with the --outdated option")
	showCmd.Flags().BoolP("installed", "i", false, "List installed packages only (enabled by default, only present for BC).")
	showCmd.Flags().BoolP("latest", "l", false, "Show the latest version")
	showCmd.Flags().Bool("locked", false, "List all locked packages")
	showCmd.Flags().BoolP("major-only", "M", false, "Show only packages that have major SemVer-compatible updates. Use with the --latest or --outdated option.")
	showCmd.Flags().BoolP("minor-only", "m", false, "Show only packages that have minor SemVer-compatible updates. Use with the --latest or --outdated option.")
	showCmd.Flags().BoolP("name-only", "N", false, "List package names only")
	showCmd.Flags().Bool("no-dev", false, "Disables search in require-dev packages.")
	showCmd.Flags().BoolP("outdated", "o", false, "Show the latest version but only for packages that are outdated")
	showCmd.Flags().Bool("patch-only", false, "Show only packages that have patch SemVer-compatible updates. Use with the --latest or --outdated option.")
	showCmd.Flags().BoolP("path", "P", false, "Show package paths")
	showCmd.Flags().BoolP("platform", "p", false, "List platform packages only")
	showCmd.Flags().BoolP("self", "s", false, "Show the root package information")
	showCmd.Flags().BoolP("sort-by-age", "A", false, "Displays the installed version's age, and sorts packages oldest first. Use with the --latest or --outdated option.")
	showCmd.Flags().Bool("strict", false, "Return a non-zero exit code when there are outdated packages")
	showCmd.Flags().BoolP("tree", "t", false, "List the dependencies as a tree")

	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"format":              carapace.ActionValues("json", "text"),
		"ignore":              action.ActionInstalledPackages(showCmd),
		"ignore-platform-req": action.ActionPlatformPackages(showCmd),
	})

	carapace.Gen(showCmd).PositionalCompletion(
		action.ActionPackagesByMode(showCmd),
	)
}
