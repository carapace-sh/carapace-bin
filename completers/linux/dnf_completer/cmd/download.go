package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [options] <package-spec>...",
	Short: "download packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()

	downloadCmd.Flags().String("arch", "", "limit to packages of given architectures")
	downloadCmd.Flags().Bool("resolve", false, "resolve dependencies and download missing ones")
	downloadCmd.Flags().Bool("alldeps", false, "download all dependencies")
	downloadCmd.Flags().String("destdir", "", "set directory for downloading packages")
	downloadCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not available")
	downloadCmd.Flags().Bool("srpm", false, "download the source rpm")
	downloadCmd.Flags().Bool("debuginfo", false, "download the debuginfo rpm")
	downloadCmd.Flags().Bool("debugsource", false, "download the debugsource rpm")
	downloadCmd.Flags().Bool("url", false, "print list of URLs where the rpms can be downloaded")
	downloadCmd.Flags().String("urlprotocol", "", "filter URLs by protocol")
	downloadCmd.Flags().Bool("allmirrors", false, "print URLs from all available mirrors")

	rootCmd.AddCommand(downloadCmd)

	carapace.Gen(downloadCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(downloadCmd),
	)
}
