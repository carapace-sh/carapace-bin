package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var balance_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve your account balance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(balance_getCmd).Standalone()
	balance_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `MonthToDateBalance`, `AccountBalance`, `MonthToDateUsage`, `GeneratedAt`")
	balance_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	balanceCmd.AddCommand(balance_getCmd)

	carapace.Gen(balance_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("MonthToDateBalance", "AccountBalance", "MonthToDateUsage", "GeneratedAt").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
