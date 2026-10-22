package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var accelerator_arq_bindCmd = &cobra.Command{
	Use:   "bind",
	Short: "Bind accelerator to instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_arq_bindCmd).Standalone()

	accelerator_arq_bindCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_arq_bindCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_arq_bindCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_arq_bindCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_arq_bindCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_arq_bindCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	accelerator_arq_bindCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_arq_bindCmd.Flags().String("variable", "", "==SUPPRESS==")
	accelerator_arqCmd.AddCommand(accelerator_arq_bindCmd)

	carapace.Gen(accelerator_arq_bindCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns(accelerator_arq_bindCmd),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})

	carapace.Gen(accelerator_arq_bindCmd).PositionalCompletion(action.ActionAcceleratorArqs(accelerator_arq_bindCmd))
}
