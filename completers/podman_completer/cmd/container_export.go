package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_exportCmd = &cobra.Command{
	Use:   "export [options] CONTAINER",
	Short: "Export container's filesystem contents as a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_exportCmd).Standalone()

	container_exportCmd.Flags().StringP("output", "o", "", "Write to a specified file (default: stdout, which must be redirected)")
	containerCmd.AddCommand(container_exportCmd)
}
