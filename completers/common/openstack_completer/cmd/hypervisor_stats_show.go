package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hypervisor_stats_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display hypervisor stats details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hypervisor_stats_showCmd).Standalone()

	hypervisor_stats_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	hypervisor_stats_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	hypervisor_stats_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	hypervisor_stats_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	hypervisor_stats_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	hypervisor_stats_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	hypervisor_stats_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	hypervisor_stats_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	hypervisor_statsCmd.AddCommand(hypervisor_stats_showCmd)
}
