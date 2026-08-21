package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mapping_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new mapping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapping_createCmd).Standalone()

	mapping_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	mapping_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	mapping_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	mapping_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	mapping_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	mapping_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	mapping_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	mapping_createCmd.Flags().String("rules", "", "Filename that contains a set of mapping rules (required)")
	mapping_createCmd.Flags().String("schema-version", "", "The federated attribute mapping schema version.")
	mapping_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	mapping_createCmd.MarkFlagRequired("rules")
	mappingCmd.AddCommand(mapping_createCmd)
}
