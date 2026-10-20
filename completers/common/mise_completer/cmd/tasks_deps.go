// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Show the dependency tree of tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_depsCmd).Standalone()

	tasks_depsCmd.Flags().Bool("compact", false, "Collapse repeated dependencies after their first occurrence")
	tasks_depsCmd.Flags().Bool("dot", false, "Display dependencies in DOT format")
	tasks_depsCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_depsCmd.Flags().Bool("hidden", false, "Show hidden tasks")
	tasksCmd.AddCommand(tasks_depsCmd)
}
