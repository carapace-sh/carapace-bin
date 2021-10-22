package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var account_ratelimitCmd = &cobra.Command{
	Use:   "ratelimit",
	Short: "Retrieve your API usage and the remaining quota",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_ratelimitCmd).Standalone()
	account_ratelimitCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Limit`, `Remaining`, `Reset`")
	account_ratelimitCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	accountCmd.AddCommand(account_ratelimitCmd)

	carapace.Gen(account_ratelimitCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Limit", "Remaining", "Reset").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

}
