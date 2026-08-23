package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setquotaCmd)
}

var setquotaCmd = &cobra.Command{
	Use:   "setquota",
	Short: "set the quota for a destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setquotaCmd).Standalone()

	carapace.Gen(setquotaCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
	)
}