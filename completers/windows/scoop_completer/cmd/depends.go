package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:   "depends",
	Short: "list dependencies for an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()
	dependsCmd.Flags().StringP("arch", "a", "", "use the specified architecture (32bit|64bit|arm64)")
	rootCmd.AddCommand(dependsCmd)

	carapace.Gen(dependsCmd).FlagCompletion(carapace.ActionMap{
		"arch": carapace.ActionValues("32bit", "64bit", "arm64"),
	})
}
