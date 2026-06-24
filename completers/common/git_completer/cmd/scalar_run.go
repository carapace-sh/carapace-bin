package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Scalar maintenance task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_runCmd).Standalone()

	scalar_runCmd.Flags().String("task", "", "run a specific task")
	scalarCmd.AddCommand(scalar_runCmd)

	carapace.Gen(scalar_runCmd).FlagCompletion(carapace.ActionMap{
		"task": carapace.ActionValues("all", "config", "commit-graph", "fetch", "loose-objects", "pack-files"),
	})
}
