package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireConfigureS3UploadCleanupCmd = &cobra.Command{
	Use:   "livewire:configure-s3-upload-cleanup",
	Short: "Configure temporary file upload s3 directory to automatically cleanup files older than 24hrs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireConfigureS3UploadCleanupCmd).Standalone()

	rootCmd.AddCommand(livewireConfigureS3UploadCleanupCmd)
}
