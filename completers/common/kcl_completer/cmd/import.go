package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "KCL import tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().BoolP("force", "f", false, "Force overwrite output file")
	importCmd.Flags().StringP("identifier", "i", "", "Assign the imported data to this top-level KCL name")
	importCmd.Flags().StringP("mode", "m", "auto", "Specify the import mode")
	importCmd.Flags().StringP("output", "o", "", "Specify the output file path")
	importCmd.Flags().StringP("package", "p", "models", "The package to save the models")
	importCmd.Flags().BoolP("skip-validation", "s", false, "Skips validation of spec prior to generation")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"mode":   carapace.ActionValues("auto", "crd", "gostruct", "json", "jsonschema", "openapi", "terraformschema", "toml", "yaml"),
		"output": carapace.ActionFiles(".k"),
	})

	carapace.Gen(importCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
