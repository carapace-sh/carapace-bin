package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ExportCmd).Standalone()
	rootCmd.AddCommand(ExportCmd)
}
