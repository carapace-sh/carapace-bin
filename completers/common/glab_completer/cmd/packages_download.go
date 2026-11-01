package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packages_downloadCmd = &cobra.Command{
	Use:     "download --name <package> --version <version> --filename <file> [flags]",
	Short:   "Download a file from a project's package registry.",
	Aliases: []string{"dl"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packages_downloadCmd).Standalone()

	packages_downloadCmd.Flags().String("filename", "", "Name of the file within the package to download.")
	packages_downloadCmd.Flags().Bool("force", false, "Overwrite the target file if it already exists.")
	packages_downloadCmd.Flags().StringP("name", "n", "", "Name of the package.")
	packages_downloadCmd.Flags().Bool("no-verify", false, "Do not verify the checksum of the downloaded file. Warning: when enabled, this setting allows the download of files that are corrupt or tampered with.")
	packages_downloadCmd.Flags().StringP("path", "p", "", "Directory to save the file in (keeps its original name) or a full file path to rename it. Defaults to the original name in the current directory.")
	packages_downloadCmd.Flags().String("version", "", "Version of the package.")
	packages_downloadCmd.MarkFlagRequired("filename")
	packages_downloadCmd.MarkFlagRequired("name")
	packages_downloadCmd.MarkFlagRequired("version")
	packagesCmd.AddCommand(packages_downloadCmd)

	carapace.Gen(packages_downloadCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})
}
