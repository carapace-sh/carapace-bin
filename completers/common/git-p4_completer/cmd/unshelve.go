package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var p4UnshelveCmd = &cobra.Command{
	Use:   "unshelve",
	Short: "Unshelve a shelved P4 changelist into a git commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(p4UnshelveCmd).Standalone()

	p4UnshelveCmd.Flags().String("origin", "", "Sets the git refspec against which the shelved P4 changelist is compared")
	rootCmd.AddCommand(p4UnshelveCmd)

	carapace.Gen(p4UnshelveCmd).FlagCompletion(carapace.ActionMap{
		"origin": carapace.ActionValues(),
	})

	carapace.Gen(p4UnshelveCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
