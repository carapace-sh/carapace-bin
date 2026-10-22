package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_rule_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display access rule details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_rule_showCmd).Standalone()

	access_rule_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	access_rule_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	access_rule_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	access_rule_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	access_rule_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	access_rule_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	access_rule_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	access_rule_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	access_ruleCmd.AddCommand(access_rule_showCmd)

	carapace.Gen(access_rule_showCmd).FlagCompletion(carapace.ActionMap{
		"column": action.ActionColumns("access", "rule", "list"),
		"format": carapace.ActionValues("json", "shell", "table", "value", "yaml"),
	})
}
