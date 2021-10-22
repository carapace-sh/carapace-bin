package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vpcs_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List VPCs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcs_listCmd).Standalone()
	vpcs_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `URN`, `Name`, `Description`, `IPRange`, `Region`, `Created`, `Default`")
	vpcs_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	vpcsCmd.AddCommand(vpcs_listCmd)
}
