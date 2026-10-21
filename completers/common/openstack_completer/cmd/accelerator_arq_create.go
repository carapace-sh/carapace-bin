package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_arq_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Register a new accelerator request with the accelerator service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_arq_createCmd).Standalone()

	accelerator_arq_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_arq_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_arq_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_arq_createCmd.Flags().String("group-id", "", "==SUPPRESS==")
	accelerator_arq_createCmd.Flags().String("image-uuid", "", "==SUPPRESS==")
	accelerator_arq_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_arq_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_arq_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	accelerator_arq_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_arq_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	accelerator_arqCmd.AddCommand(accelerator_arq_createCmd)
}
