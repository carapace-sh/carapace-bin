package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var account_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve account profile details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_getCmd).Standalone()
	account_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Email`, `DropletLimit`, `EmailVerified`, `UUID`, `Status`")
	account_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	accountCmd.AddCommand(account_getCmd)

	carapace.Gen(account_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Email", "DropletLimit", "EmailVerified", "UUID", "Status").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
