package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new network flavor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_createCmd).Standalone()

	network_flavor_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_flavor_createCmd.Flags().String("description", "", "Description for the flavor")
	network_flavor_createCmd.Flags().Bool("disable", false, "Disable the flavor")
	network_flavor_createCmd.Flags().Bool("enable", false, "Enable the flavor (default)")
	network_flavor_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_flavor_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_flavor_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_flavor_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_flavor_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_flavor_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_flavor_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_flavor_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	network_flavor_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_flavor_createCmd.Flags().String("service-type", "", "Service type to which the flavor applies.")
	network_flavor_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_flavor_createCmd.MarkFlagRequired("service-type")
	network_flavorCmd.AddCommand(network_flavor_createCmd)
}
