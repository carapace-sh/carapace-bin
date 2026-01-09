package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports configuration resources to a configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_exportCmd).Standalone()

	configure_exportCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	configure_exportCmd.Flags().Bool("all", false, "Exports all package configurations.")
	configure_exportCmd.Flags().Bool("include-versions", false, "Include package versions in export file")
	configure_exportCmd.Flags().String("module", "", "The module of the resource to export.")
	configure_exportCmd.Flags().StringP("output", "o", "", "File where the result is to be written")
	configure_exportCmd.Flags().String("package-id", "", "The package identifier to export.")
	configure_exportCmd.Flags().BoolP("recurse", "r", false, "Exports all package configurations.")
	configure_exportCmd.Flags().String("resource", "", "The configuration resource to export.")
	configure_exportCmd.Flags().StringP("source", "s", "", "Export packages from the specified source")
	configureCmd.AddCommand(configure_exportCmd)

	carapace.Gen(configure_exportCmd).FlagCompletion(carapace.ActionMap{
		"module":     carapace.ActionValues(),
		"output":     carapace.ActionFiles(),
		"package-id": carapace.ActionValues(),
		"resource":   carapace.ActionValues(),
		"source":     carapace.ActionValues(),
	})
}
