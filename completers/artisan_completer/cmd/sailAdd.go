package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sailAddCmd = &cobra.Command{
	Use:   "sail:add [<services>]",
	Short: "Add a service to an existing Sail installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sailAddCmd).Standalone()

	rootCmd.AddCommand(sailAddCmd)
}
