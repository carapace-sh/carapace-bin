package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var buildImportBundleCmd = &cobra.Command{
	Use:     "build-import-bundle [OPTION…] LOCATION FILENAME",
	Short:   "Import a file bundle into a local repository",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildImportBundleCmd).Standalone()

	buildImportBundleCmd.Flags().String("gpg-homedir", "", "GPG Homedir to use when looking for keyrings")
	buildImportBundleCmd.Flags().String("gpg-sign", "", "GPG Key ID to sign the commit with")
	buildImportBundleCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildImportBundleCmd.Flags().Bool("no-summary-index", false, "Don't generate a summary index")
	buildImportBundleCmd.Flags().Bool("no-update-summary", false, "Don't update the summary")
	buildImportBundleCmd.Flags().Bool("oci", false, "Import oci image instead of flatpak bundle")
	buildImportBundleCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildImportBundleCmd.Flags().String("ref", "", "Override the ref used for the imported bundle")
	buildImportBundleCmd.Flags().Bool("update-appstream", false, "Update the appstream branch")
	buildImportBundleCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(buildImportBundleCmd)

	carapace.Gen(buildImportBundleCmd).FlagCompletion(carapace.ActionMap{
		"gpg-homedir": carapace.ActionDirectories(),
		"gpg-sign":    os.ActionGpgKeyIds(), // TODO handle gpg-homedir flag
		// "ref":         carapace.ActionValues(), // TODO
	})

	carapace.Gen(buildImportBundleCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO location
		carapace.ActionFiles(),
	)
}
