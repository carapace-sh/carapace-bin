package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_flow_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new tap flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_flow_createCmd).Standalone()

	tap_flow_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_flow_createCmd.Flags().String("description", "", "Description of the tap flow.")
	tap_flow_createCmd.Flags().String("direction", "", "Direction of the Tap flow.")
	tap_flow_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_flow_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_flow_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_flow_createCmd.Flags().String("name", "", "Name of the tap flow.")
	tap_flow_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_flow_createCmd.Flags().String("port", "", "Source port (name or ID) to monitor.")
	tap_flow_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	tap_flow_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_flow_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	tap_flow_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	tap_flow_createCmd.Flags().String("tap-service", "", "Tap service (name or ID) to associate with this tap flow.")
	tap_flow_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	tap_flow_createCmd.Flags().String("vlan-filter", "", "VLAN IDs to mirror in the form of range string.")
	tap_flow_createCmd.MarkFlagRequired("direction")
	tap_flow_createCmd.MarkFlagRequired("port")
	tap_flow_createCmd.MarkFlagRequired("tap-service")
	tap_flowCmd.AddCommand(tap_flow_createCmd)
}
