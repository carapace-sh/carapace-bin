package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export current sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_exportCmd).Standalone()

	source_exportCmd.Flags().StringP("name", "n", "", "Name of the source")
	sourceCmd.AddCommand(source_exportCmd)

	// TODO flag completion
}
