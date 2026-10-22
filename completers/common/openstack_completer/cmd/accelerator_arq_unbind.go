package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var accelerator_arq_unbindCmd = &cobra.Command{
	Use:   "unbind",
	Short: "Unbind accelerator from instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_arq_unbindCmd).Standalone()

	accelerator_arq_unbindCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_arq_unbindCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_arq_unbindCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_arq_unbindCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_arq_unbindCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_arq_unbindCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	accelerator_arq_unbindCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_arq_unbindCmd.Flags().String("variable", "", "==SUPPRESS==")
	accelerator_arqCmd.AddCommand(accelerator_arq_unbindCmd)

	carapace.Gen(accelerator_arq_unbindCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns(accelerator_arq_unbindCmd),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})

	carapace.Gen(accelerator_arq_unbindCmd).PositionalCompletion(action.ActionAcceleratorArqs())
}
