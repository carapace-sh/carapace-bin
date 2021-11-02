package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vpcs_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a VPC",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcs_getCmd).Standalone()
	vpcs_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `URN`, `Name`, `Description`, `IPRange`, `Region`, `Created`, `Default`")
	vpcs_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	vpcsCmd.AddCommand(vpcs_getCmd)
}
