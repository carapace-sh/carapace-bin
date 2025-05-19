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

	securefile_downloadCmd.Flags().StringP("path", "p", "", "Path to download the secure file to, including filename and extension.")
	securefileCmd.AddCommand(securefile_downloadCmd)

	carapace.Gen(securefile_downloadCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})

	// TODO complete file ids
}
