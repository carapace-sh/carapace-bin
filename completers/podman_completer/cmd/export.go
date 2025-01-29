package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [options] CONTAINER",
	Short: "Export container's filesystem contents as a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().StringP("output", "o", "", "Write to a specified file (default: stdout, which must be redirected)")
	rootCmd.AddCommand(exportCmd)
}
