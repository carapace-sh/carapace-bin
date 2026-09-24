package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a network trunk for a given project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_createCmd).Standalone()

	network_trunk_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_trunk_createCmd.Flags().String("description", "", "A description of the trunk")
	network_trunk_createCmd.Flags().Bool("disable", false, "Disable trunk")
	network_trunk_createCmd.Flags().Bool("enable", false, "Enable trunk (default)")
	network_trunk_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_trunk_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_trunk_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_trunk_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_trunk_createCmd.Flags().String("parent-port", "", "Parent port belonging to this trunk (name or ID)")
	network_trunk_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_trunk_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_trunk_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	network_trunk_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_trunk_createCmd.Flags().String("subport", "", "Subport to add.")
	network_trunk_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_trunk_createCmd.MarkFlagRequired("parent-port")
	network_trunkCmd.AddCommand(network_trunk_createCmd)
}
