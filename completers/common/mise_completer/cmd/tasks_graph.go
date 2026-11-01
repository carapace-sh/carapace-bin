package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Show the monorepo project graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_graphCmd).Standalone()

	tasks_graphCmd.Flags().Bool("explain", false, "Show which provider and workspace file each project, dependency, and task field came from")
	tasks_graphCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_graphCmd.Flags().BoolP("json", "J", false, "Output the project graph as JSON")
	tasks_graphCmd.Flags().Bool("no-header", false, "Do not print table headers")
	tasks_graphCmd.Flags().Bool("no-headers", false, "Do not print table headers")
	tasks_graphCmd.Flag("no-headers").Hidden = true
	tasksCmd.AddCommand(tasks_graphCmd)
}
