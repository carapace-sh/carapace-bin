package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var limit_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display limit details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limit_showCmd).Standalone()

	limit_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	limit_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	limit_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	limit_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	limit_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	limit_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	limit_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	limit_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	limitCmd.AddCommand(limit_showCmd)

	carapace.Gen(limit_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns(limit_showCmd),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})

	carapace.Gen(limit_showCmd).PositionalCompletion(action.ActionLimits(limit_showCmd))
}
