package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports a list of the installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	exportCmd.Flags().Bool("include-versions", false, "Include package versions in export file")
	exportCmd.Flags().StringP("output", "o", "", "File where the result is to be written")
	exportCmd.Flags().StringP("source", "s", "", "Export packages from the specified source")
	rootCmd.AddCommand(exportCmd)

	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
