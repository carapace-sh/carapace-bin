package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_default_statefulness_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a default statefulness setting for security groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_default_statefulness_createCmd).Standalone()

	security_group_default_statefulness_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	security_group_default_statefulness_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	security_group_default_statefulness_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	security_group_default_statefulness_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	security_group_default_statefulness_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	security_group_default_statefulness_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	security_group_default_statefulness_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	security_group_default_statefulness_createCmd.Flags().String("project", "", "Apply the setting to this project (name or ID).")
	security_group_default_statefulness_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	security_group_default_statefulness_createCmd.Flags().Bool("stateful", false, "Set default statefulness to stateful")
	security_group_default_statefulness_createCmd.Flags().Bool("stateless", false, "Set default statefulness to stateless")
	security_group_default_statefulness_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	security_group_default_statefulnessCmd.AddCommand(security_group_default_statefulness_createCmd)
}
