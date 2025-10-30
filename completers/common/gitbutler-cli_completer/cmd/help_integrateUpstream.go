package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_integrateUpstreamCmd = &cobra.Command{
	Use:   "integrate-upstream",
	Short: "Update the local workspace against an updated remote or target branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_integrateUpstreamCmd).Standalone()

	helpCmd.AddCommand(help_integrateUpstreamCmd)
}
