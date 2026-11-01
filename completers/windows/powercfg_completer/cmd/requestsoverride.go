package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var requestsoverrideCmd = &cobra.Command{
	Use:   "requestsoverride",
	Short: "set a power request override for a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(requestsoverrideCmd).Standalone()
	rootCmd.AddCommand(requestsoverrideCmd)
}
