package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printTokenCmd = &cobra.Command{
	Use:    "print-token",
	Short:  "Pretty-print the contents of a JWT token",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printTokenCmd).Standalone()

	rootCmd.AddCommand(printTokenCmd)

	carapace.Gen(printTokenCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
