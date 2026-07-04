package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Compute checksums of distribution files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checksumCmd).Standalone()
	rootCmd.AddCommand(checksumCmd)
}
