package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qfailureflagCmd = &cobra.Command{
	Use:   "qfailureflag",
	Short: "query the failure flag of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qfailureflagCmd).Standalone()
	rootCmd.AddCommand(qfailureflagCmd)
}
