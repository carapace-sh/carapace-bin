package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var leavesCmd = &cobra.Command{
	Use:   "leaves",
	Short: "list groups of leaf packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(leavesCmd).Standalone()

	rootCmd.AddCommand(leavesCmd)
}
