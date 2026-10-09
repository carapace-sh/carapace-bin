package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new firewall group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_createCmd).Standalone()

	firewall_group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	firewall_group_createCmd.Flags().String("description", "", "Description of the firewall group")
	firewall_group_createCmd.Flags().Bool("disable", false, "Disable firewall group")
	firewall_group_createCmd.Flags().String("egress-firewall-policy", "", "Egress firewall policy (name or ID)")
	firewall_group_createCmd.Flags().Bool("enable", false, "Enable firewall group")
	firewall_group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	firewall_group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	firewall_group_createCmd.Flags().String("ingress-firewall-policy", "", "Ingress firewall policy (name or ID)")
	firewall_group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	firewall_group_createCmd.Flags().String("name", "", "(Deprecated, please pass name as a positional argument) Name for the firewall group")
	firewall_group_createCmd.Flags().Bool("no-egress-firewall-policy", false, "Detach egress firewall policy from the firewall group")
	firewall_group_createCmd.Flags().Bool("no-ingress-firewall-policy", false, "Detach ingress firewall policy from the firewall group")
	firewall_group_createCmd.Flags().Bool("no-port", false, "Detach all port from the firewall group")
	firewall_group_createCmd.Flags().Bool("no-share", false, "Restrict use of the firewall group to the current project")
	firewall_group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	firewall_group_createCmd.Flags().String("port", "", "Port(s) (name or ID) to apply firewall group.")
	firewall_group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	firewall_group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	firewall_group_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	firewall_group_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	firewall_group_createCmd.Flags().Bool("share", false, "Share the firewall group to be used in all projects (by default, it is restricted to be used by the current project).")
	firewall_group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	firewall_groupCmd.AddCommand(firewall_group_createCmd)
}
