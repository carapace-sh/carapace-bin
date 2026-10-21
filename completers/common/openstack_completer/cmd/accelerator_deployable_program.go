package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_deployable_programCmd = &cobra.Command{
	Use:   "program",
	Short: "Reconfigure deployable with a new bitstream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_deployable_programCmd).Standalone()

	accelerator_deployable_programCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_deployable_programCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_deployable_programCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_deployable_programCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_deployable_programCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_deployable_programCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	accelerator_deployable_programCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_deployable_programCmd.Flags().String("variable", "", "==SUPPRESS==")
	accelerator_deployableCmd.AddCommand(accelerator_deployable_programCmd)
}
