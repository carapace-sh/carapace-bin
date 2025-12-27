package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:     "verify",
	Short:   "Verify package integrity and signature",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "repository maintenance",
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	rootCmd.AddCommand(verifyCmd)

	carapace.Gen(verifyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
