package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a QoS policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_policy_createCmd).Standalone()

	network_qos_policy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_qos_policy_createCmd.Flags().Bool("default", false, "Set this as a default network QoS policy")
	network_qos_policy_createCmd.Flags().String("description", "", "Description of the QoS policy")
	network_qos_policy_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_qos_policy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_qos_policy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_qos_policy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_qos_policy_createCmd.Flags().Bool("no-default", false, "Set this as a non-default network QoS policy")
	network_qos_policy_createCmd.Flags().Bool("no-share", false, "Make the QoS policy not accessible by other projects (default)")
	network_qos_policy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_qos_policy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_qos_policy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_qos_policy_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	network_qos_policy_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_qos_policy_createCmd.Flags().Bool("share", false, "Make the QoS policy accessible by other projects")
	network_qos_policy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_qos_policyCmd.AddCommand(network_qos_policy_createCmd)
}
