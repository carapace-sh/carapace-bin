package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Download a bottle (if available) or source packages for <formula>e and binaries for <cask>s",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().Bool("HEAD", false, "Fetch HEAD version instead of stable version.")
	fetchCmd.Flags().Bool("arch", false, "Download for the given CPU architecture. (Pass `all` to download for all architectures.)")
	fetchCmd.Flags().Bool("bottle-tag", false, "Download a bottle for given tag.")
	fetchCmd.Flags().Bool("build-bottle", false, "Download source packages (for eventual bottling) rather than a bottle.")
	fetchCmd.Flags().Bool("build-from-source", false, "Download source packages rather than a bottle.")
	fetchCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	fetchCmd.Flags().Bool("debug", false, "Display any debugging information.")
	fetchCmd.Flags().Bool("deps", false, "Also download dependencies for any listed <formula>.")
	fetchCmd.Flags().Bool("force", false, "Remove a previously cached version and re-fetch.")
	fetchCmd.Flags().Bool("force-bottle", false, "Download a bottle if it exists for the current or newest version of macOS, even if it would not be used during installation.")
	fetchCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	fetchCmd.Flags().Bool("help", false, "Show this message.")
	fetchCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	fetchCmd.Flags().Bool("os", false, "Download for the given operating system. (Pass `all` to download for all operating systems.)")
	fetchCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	fetchCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	fetchCmd.Flags().Bool("retry", false, "Retry if downloading fails or re-download if the checksum of a previously cached version no longer matches. Tries at most 5 times with exponential backoff.")
	fetchCmd.Flags().Bool("verbose", false, "Do a verbose VCS checkout, if the URL represents a VCS. This is useful for seeing if an existing VCS cache has been updated.")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		action.ActionSearch(fetchCmd),
	)
}
