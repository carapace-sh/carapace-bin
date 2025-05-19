package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitCodesCmd = &cobra.Command{
	Use:    "exit-codes",
	Short:  "Exit codes used by gh",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitCodesCmd).Standalone()

	rootCmd.AddCommand(exitCodesCmd)
}
