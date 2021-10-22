package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Display commands that manage infrastructure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeCmd).Standalone()
	rootCmd.AddCommand(computeCmd)
}
