package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_exportCmd = &cobra.Command{
	Use:   "export [options] VOLUME",
	Short: "Export volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_exportCmd).Standalone()

	volume_exportCmd.Flags().StringP("output", "o", "", "Write to a specified file (default: stdout, which must be redirected)")
	volumeCmd.AddCommand(volume_exportCmd)
}
