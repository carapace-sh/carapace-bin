package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var failureflagCmd = &cobra.Command{
	Use:   "failureflag",
	Short: "set the failure flag for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(failureflagCmd).Standalone()
	rootCmd.AddCommand(failureflagCmd)
}
