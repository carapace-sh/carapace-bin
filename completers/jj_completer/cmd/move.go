package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move changes from one revision into another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	moveCmd.Flags().String("from", "", "Move part of this change into the destination")
	moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	moveCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to move")
	moveCmd.Flags().String("to", "", "Move part of the source into this change")
	rootCmd.AddCommand(moveCmd)
}
