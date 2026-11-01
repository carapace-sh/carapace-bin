package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recursiveCmd = &cobra.Command{
	Use:     "recursive",
	Short:   "Concurrently runs a command in all subdirectory projects",
	Aliases: []string{"multi", "m"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recursiveCmd).Standalone()

	recursiveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(recursiveCmd)
}
