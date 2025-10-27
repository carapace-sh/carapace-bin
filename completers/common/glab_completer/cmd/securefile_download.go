package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_downloadCmd = &cobra.Command{
	Use:   "download <fileID> [flags]",
	Short: "Download a secure file for a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_downloadCmd).Standalone()

	securefile_downloadCmd.Flags().Bool("all", false, "Download all (limit 100) of a project's secure files. Files are downloaded with their original name and file extension.")
	securefile_downloadCmd.Flags().Bool("force-download", false, "Force download file(s) even if checksum verification fails. Warning: when enabled, this setting allows the download of files that are corrupt or tampered with.")
	securefile_downloadCmd.Flags().Bool("no-verify", false, "Do not verify the checksum of the downloaded file(s). Warning: when enabled, this setting allows the download of files that are corrupt or tampered with.")
	securefile_downloadCmd.Flags().String("output-dir", "", "Output directory for files downloaded with --all.")
	securefile_downloadCmd.Flags().StringP("path", "p", "", "Path to download the secure file to, including filename and extension.")
	securefileCmd.AddCommand(securefile_downloadCmd)

	carapace.Gen(securefile_downloadCmd).FlagCompletion(carapace.ActionMap{
		"output-dir": carapace.ActionDirectories(),
		"path":       carapace.ActionFiles(),
	})

	// TODO complete file ids
}
