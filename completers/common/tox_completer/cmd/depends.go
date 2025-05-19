package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:     "depends",
	Aliases: []string{"de"},
	Short:   "visualize tox environment dependencies",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()

	addCommonSubcommandFlags(dependsCmd)
	rootCmd.AddCommand(dependsCmd)
}
