package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var experimental_generate_debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Shows generate debug information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_generate_debugCmd).Standalone()

	experimental_generateCmd.AddCommand(experimental_generate_debugCmd)
}
