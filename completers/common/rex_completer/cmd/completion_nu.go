package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionNuCmd = &cobra.Command{
	Use:   "nu",
	Short: "Generate Nushell completion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionNuCmd).Standalone()

	completionCmd.AddCommand(completionNuCmd)
}
