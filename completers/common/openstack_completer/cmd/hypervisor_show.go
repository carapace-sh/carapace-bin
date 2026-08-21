package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hypervisor_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display hypervisor details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hypervisor_showCmd).Standalone()

	hypervisor_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	hypervisor_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	hypervisor_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	hypervisor_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	hypervisor_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	hypervisor_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	hypervisor_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	hypervisor_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	hypervisorCmd.AddCommand(hypervisor_showCmd)
}
