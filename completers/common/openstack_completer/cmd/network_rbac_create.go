package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_rbac_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create network RBAC policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rbac_createCmd).Standalone()

	network_rbac_createCmd.Flags().String("action", "", "Action for the RBAC policy (\"access_as_external\" or \"access_as_shared\")")
	network_rbac_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_rbac_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_rbac_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_rbac_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_rbac_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_rbac_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_rbac_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_rbac_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_rbac_createCmd.Flags().String("project", "", "The owner project (name or ID)")
	network_rbac_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_rbac_createCmd.Flags().Bool("target-all-projects", false, "Allow creating RBAC policy for all projects")
	network_rbac_createCmd.Flags().String("target-project", "", "The project to which the RBAC policy will be enforced (name or ID)")
	network_rbac_createCmd.Flags().String("target-project-domain", "", "Domain the target project belongs to (name or ID).")
	network_rbac_createCmd.Flags().String("type", "", "Type of the object that RBAC policy affects (\"address_group\", \"address_scope\", \"security_group\", \"subnetpool\", \"qos_policy\" or \"network\")")
	network_rbac_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_rbac_createCmd.MarkFlagRequired("action")
	network_rbac_createCmd.MarkFlagRequired("type")
	network_rbacCmd.AddCommand(network_rbac_createCmd)
}
