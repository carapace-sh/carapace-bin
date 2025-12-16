package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lsns_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsns",
	Short: "List system namespaces",
	Long:  "https://man7.org/linux/man-pages/man8/lsns.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("filter", "Q", "", "apply display filter")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("list", "l", false, "use list format output")
	rootCmd.Flags().BoolP("list-columns", "H", false, "list the available columns")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().BoolP("notruncate", "u", false, "don't truncate text in columns")
	rootCmd.Flags().BoolP("nowrap", "W", false, "don't use multi-line representation")
	rootCmd.Flags().StringP("output", "o", "", "define which output columns to use")
	rootCmd.Flags().Bool("output-all", false, "output all columns")
	rootCmd.Flags().BoolP("persistent", "P", false, "namespaces without processes")
	rootCmd.Flags().BoolP("raw", "r", false, "use the raw output format")
	rootCmd.Flags().StringP("task", "p", "", "print process namespaces")
	rootCmd.Flags().StringP("tree", "T", "", "use tree format (parent, owner, or process)")
	rootCmd.Flags().StringP("type", "t", "", "namespace type (mnt, net, ipc, user, pid, uts, cgroup, time)")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("tree").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": action.ActionColumns().UniqueList(","),
		"task":   ps.ActionProcessIds(),
		"tree":   carapace.ActionValues("parent", "owner", "process"),
		"type":   carapace.ActionValues("mnt", "net", "ipc", "user", "pid", "uts", "cgroup", "time"),
	})
}
