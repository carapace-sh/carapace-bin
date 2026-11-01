package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Shows a list of installed packages that have updates available, including their latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().BoolP("all", "a", false, "Show all installed packages with their latest versions")
	outdatedCmd.Flags().BoolP("direct", "D", false, "Shows only packages that are directly required by the root package")
	outdatedCmd.Flags().StringP("format", "f", "", "Format of the output: text or json")
	outdatedCmd.Flags().StringArray("ignore", nil, "Ignore specified package(s). Can contain wildcards (*). Use it if you don't want to be informed about new versions of some packages.")
	outdatedCmd.Flags().StringArray("ignore-platform-req", nil, "Ignore a specific platform requirement (php & ext- packages). Use with the --outdated option")
	outdatedCmd.Flags().Bool("ignore-platform-reqs", false, "Ignore all platform requirements (php & ext- packages). Use with the --outdated option")
	outdatedCmd.Flags().Bool("locked", false, "Shows updates for packages from the lock file, regardless of what is currently in vendor dir")
	outdatedCmd.Flags().BoolP("major-only", "M", false, "Show only packages that have major SemVer-compatible updates.")
	outdatedCmd.Flags().BoolP("minor-only", "m", false, "Show only packages that have minor SemVer-compatible updates.")
	outdatedCmd.Flags().Bool("no-dev", false, "Disables search in require-dev packages.")
	outdatedCmd.Flags().BoolP("outdated", "o", false, "Show only packages that are outdated (this is the default, but present here for compat with `show`")
	outdatedCmd.Flags().BoolP("patch-only", "p", false, "Show only packages that have patch SemVer-compatible updates.")
	outdatedCmd.Flags().BoolP("sort-by-age", "A", false, "Displays the installed version's age, and sorts packages oldest first.")
	outdatedCmd.Flags().Bool("strict", false, "Return a non-zero exit code when there are outdated packages")

	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).FlagCompletion(carapace.ActionMap{
		"format":              carapace.ActionValues("json", "text"),
		"ignore":              action.ActionInstalledPackages(outdatedCmd),
		"ignore-platform-req": action.ActionPlatformPackages(outdatedCmd),
	})

	carapace.Gen(outdatedCmd).PositionalCompletion(
		action.ActionInstalledPackages(outdatedCmd),
	)
}
