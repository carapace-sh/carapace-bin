package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependencyCmd = &cobra.Command{
	Use:     "dependency",
	Short:   "manage a chart's dependencies",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyCmd).Standalone()
	rootCmd.AddCommand(dependencyCmd)
}
