package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stateCmd = &cobra.Command{
	Use:   "state <subcommand> [options] [args]",
	Short: "Advanced state management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stateCmd).Standalone()

	rootCmd.AddCommand(stateCmd)
}
