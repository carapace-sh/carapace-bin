package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_group_createCmd).Standalone()

	server_group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_group_createCmd.Flags().String("policy", "", "Add a policy to <name> Specify --os-compute-api-version 2.15 or higher for the 'soft-affinity' or 'soft-anti-affinity' policy.")
	server_group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_group_createCmd.Flags().String("rule", "", "A rule for the policy.")
	server_group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_groupCmd.AddCommand(server_group_createCmd)
}
