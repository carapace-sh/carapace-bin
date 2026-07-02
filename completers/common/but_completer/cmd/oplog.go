package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oplogCmd = &cobra.Command{
	Use:     "oplog",
	Short:   "Commands for viewing and managing operation history",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "operation history",
}

func init() {
	carapace.Gen(oplogCmd).Standalone()

	oplogCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(oplogCmd)
}
