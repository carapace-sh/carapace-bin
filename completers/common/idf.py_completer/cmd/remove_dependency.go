package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

var removeDependencyCmd = &cobra.Command{
	Use:   "remove-dependency",
	Short: "Remove specified dependency from all manifest files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeDependencyCmd).Standalone()
	rootCmd.AddCommand(removeDependencyCmd)
	carapace.Gen(removeDependencyCmd).PositionalCompletion(
		espidf.ActionDependencies(),
	)
}
