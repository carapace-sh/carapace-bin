package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oplogCmd = &cobra.Command{
	Use:   "oplog",
	Short: "Show operation history (last 20 entries)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oplogCmd).Standalone()

	oplogCmd.Flags().BoolP("help", "h", false, "Print help")
	oplogCmd.Flags().String("since", "", "Start from this oplog SHA instead of the head")
	rootCmd.AddCommand(oplogCmd)
}
