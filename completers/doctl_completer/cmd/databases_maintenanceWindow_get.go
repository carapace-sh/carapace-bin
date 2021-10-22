package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_maintenanceWindow_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details about a database cluster's maintenance windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_maintenanceWindow_getCmd).Standalone()
	databases_maintenanceWindow_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Day`, `Hour`, `Pending`")
	databases_maintenanceWindow_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_maintenanceWindowCmd.AddCommand(databases_maintenanceWindow_getCmd)

	carapace.Gen(databases_maintenanceWindow_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Day", "Hour", "Pending").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
