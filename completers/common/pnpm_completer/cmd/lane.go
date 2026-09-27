package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var laneCmd = &cobra.Command{
	Use:   "lane",
	Short: "Manage per-package release lanes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(laneCmd).Standalone()

	laneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(laneCmd)
}
