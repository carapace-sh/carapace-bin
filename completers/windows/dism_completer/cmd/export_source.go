package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ExportSourceCmd = &cobra.Command{
	Use:   "Export-Source",
	Short: "export a set of capabilities into a new repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ExportSourceCmd).Standalone()
	rootCmd.AddCommand(ExportSourceCmd)
}
