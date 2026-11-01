package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:   "depends",
	Short: "list dependencies for an app, in the order they'll be installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()
	dependsCmd.Flags().StringP("arch", "a", "", "use the specified architecture (32bit|64bit|arm64)")
	rootCmd.AddCommand(dependsCmd)

	carapace.Gen(dependsCmd).FlagCompletion(carapace.ActionMap{
		"arch": carapace.ActionValues("32bit", "64bit", "arm64"),
	})

	carapace.Gen(dependsCmd).PositionalCompletion(
		scoop.ActionAvailableApps(),
	)
}
