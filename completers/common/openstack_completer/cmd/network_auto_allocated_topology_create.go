package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_auto_allocated_topology_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the  auto allocated topology for project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_auto_allocated_topology_createCmd).Standalone()

	network_auto_allocated_topology_createCmd.Flags().Bool("check-resources", false, "Validate the requirements for auto allocated topology.")
	network_auto_allocated_topology_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_auto_allocated_topology_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_auto_allocated_topology_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_auto_allocated_topology_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_auto_allocated_topology_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_auto_allocated_topology_createCmd.Flags().Bool("or-show", false, "If topology exists returns the topology's information (default)")
	network_auto_allocated_topology_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_auto_allocated_topology_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_auto_allocated_topology_createCmd.Flags().String("project", "", "Return the auto allocated topology for a given project.")
	network_auto_allocated_topology_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_auto_allocated_topology_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_auto_allocated_topologyCmd.AddCommand(network_auto_allocated_topology_createCmd)
}
