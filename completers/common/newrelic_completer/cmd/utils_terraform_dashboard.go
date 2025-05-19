package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utils_terraform_dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Generate HCL for the newrelic_one_dashboard resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utils_terraform_dashboardCmd).Standalone()
	utils_terraform_dashboardCmd.Flags().StringP("file", "f", "", "a file that contains exported dashboard JSON")
	utils_terraform_dashboardCmd.Flags().StringP("label", "l", "", "the resource label to use when generating resource HCL")
	utils_terraform_dashboardCmd.Flags().StringP("out", "o", "", "the file to send the generated HCL to")
	utils_terraform_dashboardCmd.Flags().IntP("shiftWidth", "w", 2, "the indentation shift with of the output")
	utils_terraformCmd.AddCommand(utils_terraform_dashboardCmd)

	carapace.Gen(utils_terraform_dashboardCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"out":  carapace.ActionFiles(),
	})
}
