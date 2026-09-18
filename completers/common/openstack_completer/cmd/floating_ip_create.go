package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create floating IP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_createCmd).Standalone()

	floating_ip_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	floating_ip_createCmd.Flags().String("description", "", "Set floating IP description")
	floating_ip_createCmd.Flags().String("dns-domain", "", "Set DNS domain for this floating IP")
	floating_ip_createCmd.Flags().String("dns-name", "", "Set DNS name for this floating IP")
	floating_ip_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	floating_ip_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	floating_ip_createCmd.Flags().String("fixed-ip-address", "", "Fixed IP address mapped to the floating IP")
	floating_ip_createCmd.Flags().String("floating-ip-address", "", "Floating IP address")
	floating_ip_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	floating_ip_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	floating_ip_createCmd.Flags().Bool("no-tag", false, "No tags associated with the floating IP")
	floating_ip_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	floating_ip_createCmd.Flags().String("port", "", "Port to be associated with the floating IP (name or ID)")
	floating_ip_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	floating_ip_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	floating_ip_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	floating_ip_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	floating_ip_createCmd.Flags().String("qos-policy", "", "Attach QoS policy to the floating IP (name or ID)")
	floating_ip_createCmd.Flags().String("subnet", "", "Subnet on which you want to create the floating IP (name or ID)")
	floating_ip_createCmd.Flags().String("tag", "", "Tag to be added to the floating IP (repeat option to set multiple tags)")
	floating_ip_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	floating_ipCmd.AddCommand(floating_ip_createCmd)
}
