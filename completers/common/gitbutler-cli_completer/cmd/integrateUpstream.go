package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integrateUpstreamCmd = &cobra.Command{
	Use:     "integrate-upstream",
	Short:   "Update the local workspace against an updated remote or target branch",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integrateUpstreamCmd).Standalone()

	integrateUpstreamCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(integrateUpstreamCmd)
}
