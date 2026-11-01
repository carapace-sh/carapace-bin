package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:     "tasks",
	Short:   "Manage tasks",
	Aliases: []string{"t", "task"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasksCmd).Standalone()

	tasksCmd.Flags().BoolP("all", "a", false, "Include tasks from every project in the monorepo, not only the current directory's hierarchy")
	tasksCmd.Flags().Bool("complete", false, "Display tasks for usage completion")
	tasksCmd.Flags().BoolP("extended", "x", false, "Show all columns")
	tasksCmd.Flags().BoolP("global", "g", false, "Only show global tasks")
	tasksCmd.Flags().BoolP("help", "h", false, "Print help")
	tasksCmd.Flags().Bool("hidden", false, "Show hidden tasks")
	tasksCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	tasksCmd.Flags().BoolP("local", "l", false, "Only show non-global tasks")
	tasksCmd.Flags().Bool("name-only", false, "Print only task names, one per line (for fzf and similar tools)")
	tasksCmd.Flags().Bool("no-header", false, "Do not print the table header")
	tasksCmd.Flags().Bool("no-headers", false, "Do not print the table header")
	tasksCmd.Flags().String("sort", "", "Column to sort by (default: name)")
	tasksCmd.Flags().String("sort-order", "", "Sort direction (default: asc)")
	tasksCmd.Flags().Bool("usage", false, "")
	tasksCmd.Flag("complete").Hidden = true
	tasksCmd.Flag("no-headers").Hidden = true
	tasksCmd.Flag("usage").Hidden = true
	rootCmd.AddCommand(tasksCmd)

	carapace.Gen(tasksCmd).FlagCompletion(carapace.ActionMap{
		"sort":       carapace.ActionValues("alias", "description", "name", "source"),
		"sort-order": carapace.ActionValues("asc", "desc"),
	})

	carapace.Gen(tasksCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
