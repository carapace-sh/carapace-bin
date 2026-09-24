package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flavor_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new flavor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavor_createCmd).Standalone()

	flavor_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	flavor_createCmd.Flags().String("description", "", "Description for the flavor.(Supported")
	flavor_createCmd.Flags().String("disk", "", "Disk size in GB (default 0G)")
	flavor_createCmd.Flags().String("ephemeral", "", "Ephemeral disk size in GB (default 0G)")
	flavor_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	flavor_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	flavor_createCmd.Flags().String("id", "", "Unique flavor ID")
	flavor_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	flavor_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	flavor_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	flavor_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	flavor_createCmd.Flags().Bool("private", false, "Flavor is not available to other projects")
	flavor_createCmd.Flags().String("project", "", "Allow <project> to access private flavor (name or ID) (Must be used with --private option)")
	flavor_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	flavor_createCmd.Flags().String("property", "", "Property to add for this flavor (repeat option to set multiple properties)")
	flavor_createCmd.Flags().Bool("public", false, "Flavor is available to other projects (default)")
	flavor_createCmd.Flags().String("ram", "", "Memory size in MB (default 256M)")
	flavor_createCmd.Flags().String("rxtx-factor", "", "RX/TX factor")
	flavor_createCmd.Flags().String("swap", "", "Additional swap space size in MB (default 0M)")
	flavor_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	flavor_createCmd.Flags().String("vcpus", "", "Number of vcpus (default 1)")
	flavorCmd.AddCommand(flavor_createCmd)
}
