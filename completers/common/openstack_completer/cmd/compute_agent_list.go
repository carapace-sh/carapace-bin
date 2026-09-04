package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compute_agent_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List compute agents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_agent_listCmd).Standalone()

	compute_agent_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	compute_agent_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	compute_agent_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	compute_agent_listCmd.Flags().String("hypervisor", "", "Type of hypervisor")
	compute_agent_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	compute_agent_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	compute_agent_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	compute_agent_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	compute_agent_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	compute_agent_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	compute_agent_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	compute_agentCmd.AddCommand(compute_agent_listCmd)
}
