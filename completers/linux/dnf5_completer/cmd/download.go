package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [options] <package-spec>...",
	Short: "download software to the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()

	downloadCmd.Flags().Bool("alldeps", false, "Download all dependencies (with --resolve)")
	downloadCmd.Flags().Bool("allmirrors", false, "Print URLs from all available mirrors (with --url)")
	downloadCmd.Flags().String("arch", "", "Limit to packages of given architectures")
	downloadCmd.Flags().Bool("debuginfo", false, "Download the -debuginfo package instead")
	downloadCmd.Flags().Bool("debugsource", false, "Download the -debugsource package instead")
	downloadCmd.Flags().String("destdir", "", "Set directory used for downloading packages to")
	downloadCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	downloadCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	downloadCmd.Flags().Bool("resolve", false, "Resolve and download needed dependencies")
	downloadCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	downloadCmd.Flags().Bool("srpm", false, "Download the src.rpm instead")
	downloadCmd.Flags().Bool("url", false, "Print a URL where the rpms can be downloaded instead of downloading")
	downloadCmd.Flags().String("urlprotocol", "", "Limit to specific protocols (with --url)")

	rootCmd.AddCommand(downloadCmd)

	carapace.Gen(downloadCmd).FlagCompletion(carapace.ActionMap{
		"destdir":     carapace.ActionDirectories(),
		"urlprotocol": carapace.ActionValues("http", "https", "ftp", "file"),
	})

	carapace.Gen(downloadCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(downloadCmd),
	)
}
