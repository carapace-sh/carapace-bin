package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info [OPTION…] NAME [BRANCH]",
	Short:   "Get info about an installed app or runtime",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().String("arch", "", "Arch to use")
	infoCmd.Flags().String("file-access", "", "Query file access")
	infoCmd.Flags().BoolP("help", "h", false, "Show help options")
	infoCmd.Flags().String("installation", "", "Show specific system-wide installations")
	infoCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	infoCmd.Flags().BoolP("show-commit", "c", false, "Show commit")
	infoCmd.Flags().BoolP("show-extensions", "e", false, "Show extensions")
	infoCmd.Flags().BoolP("show-location", "l", false, "Show location")
	infoCmd.Flags().BoolP("show-metadata", "m", false, "Show metadata")
	infoCmd.Flags().BoolP("show-origin", "o", false, "Show origin")
	infoCmd.Flags().BoolP("show-permissions", "M", false, "Show permissions")
	infoCmd.Flags().BoolP("show-ref", "r", false, "Show ref")
	infoCmd.Flags().Bool("show-runtime", false, "Show runtime")
	infoCmd.Flags().Bool("show-sdk", false, "Show sdk")
	infoCmd.Flags().BoolP("show-size", "s", false, "Show size")
	infoCmd.Flags().Bool("system", false, "Show system-wide installations")
	infoCmd.Flags().Bool("user", false, "Show user installations")
	infoCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(infoCmd)

	// TODO flag
	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "file-access":  carapace.ActionValues(),
		// "installation": carapace.ActionValues(),
	})

	carapace.Gen(infoCmd).PositionalCompletion(
		flatpak.ActionApplications(), // TODO runtimes??
		// TODO branches
	)
}
