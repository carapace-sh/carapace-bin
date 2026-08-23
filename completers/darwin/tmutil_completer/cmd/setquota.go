package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setquotaCmd = &cobra.Command{
	Use:   "setquota",
	Short: "set the quota for a destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setquotaCmd).Standalone()
	rootCmd.AddCommand(setquotaCmd)

	carapace.Gen(setquotaCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
	)
}