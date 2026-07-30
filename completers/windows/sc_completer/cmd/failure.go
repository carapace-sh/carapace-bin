package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var failureCmd = &cobra.Command{
	Use:   "failure",
	Short: "set actions to take if a service fails",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(failureCmd).Standalone()
	rootCmd.AddCommand(failureCmd)
}
