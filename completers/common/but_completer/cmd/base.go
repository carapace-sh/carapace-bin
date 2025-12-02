package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var baseCmd = &cobra.Command{
	Use:     "base",
	Short:   "Commands for managing the base target branch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(baseCmd).Standalone()

	baseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(baseCmd)
}
