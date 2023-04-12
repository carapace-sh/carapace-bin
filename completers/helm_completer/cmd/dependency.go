package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "manage a chart's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyCmd).Standalone()
	rootCmd.AddCommand(dependencyCmd)
}
