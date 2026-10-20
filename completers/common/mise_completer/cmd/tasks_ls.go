package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_lsCmd).Standalone()

	tasks_lsCmd.Flags().BoolP("all", "a", false, "Include tasks from every project in the monorepo, not only the current directory's hierarchy")
	tasks_lsCmd.Flags().Bool("complete", false, "Display tasks for usage completion")
	tasks_lsCmd.Flags().BoolP("extended", "x", false, "Show all columns")
	tasks_lsCmd.Flags().BoolP("global", "g", false, "Only show global tasks")
	tasks_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_lsCmd.Flags().Bool("hidden", false, "Show hidden tasks")
	tasks_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	tasks_lsCmd.Flags().BoolP("local", "l", false, "Only show non-global tasks")
	tasks_lsCmd.Flags().Bool("name-only", false, "Print only task names, one per line (for fzf and similar tools)")
	tasks_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	tasks_lsCmd.Flags().Bool("no-headers", false, "Do not print the table header")
	tasks_lsCmd.Flags().String("sort", "", "Column to sort by (default: name)")
	tasks_lsCmd.Flags().String("sort-order", "", "Sort direction (default: asc)")
	tasks_lsCmd.Flags().Bool("usage", false, "")
	tasks_lsCmd.Flag("complete").Hidden = true
	tasks_lsCmd.Flag("no-headers").Hidden = true
	tasks_lsCmd.Flag("usage").Hidden = true
	tasksCmd.AddCommand(tasks_lsCmd)

	carapace.Gen(tasks_lsCmd).FlagCompletion(carapace.ActionMap{
		"sort":       carapace.ActionValues("name", "alias", "description", "source"),
		"sort-order": carapace.ActionValues("asc", "desc"),
	})
}
