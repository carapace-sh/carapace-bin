package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_profile_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new network flavor profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_profile_createCmd).Standalone()

	network_flavor_profile_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_flavor_profile_createCmd.Flags().String("description", "", "Description for the flavor profile")
	network_flavor_profile_createCmd.Flags().Bool("disable", false, "Disable the flavor profile")
	network_flavor_profile_createCmd.Flags().String("driver", "", "Python module path to driver.")
	network_flavor_profile_createCmd.Flags().Bool("enable", false, "Enable the flavor profile")
	network_flavor_profile_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_flavor_profile_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_flavor_profile_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_flavor_profile_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_flavor_profile_createCmd.Flags().String("metainfo", "", "Metainfo for the flavor profile.")
	network_flavor_profile_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_flavor_profile_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_flavor_profile_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_flavor_profile_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_flavor_profileCmd.AddCommand(network_flavor_profile_createCmd)
}
