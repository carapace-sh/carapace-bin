package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vetCmd = &cobra.Command{
	Use:   "vet",
	Short: "KCL validation tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vetCmd).Standalone()

	vetCmd.Flags().StringP("attribute_name", "a", "", "Specify the validate schema attribute name")
	vetCmd.Flags().StringArrayP("external", "E", []string{}, "Specify the mapping of package name and path where the package is located")
	vetCmd.Flags().String("format", "", "Validate data file format (json, yaml and toml)")
	vetCmd.Flags().String("output", "text", "Specify the validation result output format (text and json)")
	vetCmd.Flags().StringP("schema", "s", "", "Specify the validate schema name")
	rootCmd.AddCommand(vetCmd)

	carapace.Gen(vetCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "toml", "yaml"),
		"output": carapace.ActionValues("json", "text"),
	})

	carapace.Gen(vetCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".yaml", ".yml", ".toml"),
		carapace.ActionFiles(".k"),
	)
}
