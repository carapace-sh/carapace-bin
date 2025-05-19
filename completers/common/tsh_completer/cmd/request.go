package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:     "request",
	Short:   "Manage access requests.",
	Aliases: []string{"requests"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(requestCmd).Standalone()

	rootCmd.AddCommand(requestCmd)
}
