package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compute_agent_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create compute agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_agent_createCmd).Standalone()

	compute_agent_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	compute_agent_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	compute_agent_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	compute_agent_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	compute_agent_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	compute_agent_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	compute_agent_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	compute_agent_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	compute_agentCmd.AddCommand(compute_agent_createCmd)
}
