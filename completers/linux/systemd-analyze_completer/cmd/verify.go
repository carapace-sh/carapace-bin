package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check unit files for correctness",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	rootCmd.AddCommand(verifyCmd)

	carapace.Gen(verifyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
