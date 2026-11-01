package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var accelerator_deployable_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show deployable details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_deployable_showCmd).Standalone()

	accelerator_deployable_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	accelerator_deployable_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	accelerator_deployable_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	accelerator_deployable_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	accelerator_deployable_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	accelerator_deployable_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	accelerator_deployable_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	accelerator_deployable_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	accelerator_deployableCmd.AddCommand(accelerator_deployable_showCmd)

	carapace.Gen(accelerator_deployable_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns(accelerator_deployable_showCmd),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})

	carapace.Gen(accelerator_deployable_showCmd).PositionalCompletion(action.ActionAcceleratorDeployables(accelerator_deployable_showCmd))
}
