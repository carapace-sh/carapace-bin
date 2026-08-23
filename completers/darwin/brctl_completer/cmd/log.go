package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "manage CloudDocs logging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().Bool("color", false, "turn on or off color use")
	logCmd.Flags().Bool("digest", false, "only print digest logs")
	logCmd.Flags().String("filter", "", "only show lines matching predicate")
	logCmd.Flags().String("home", "", "use this as the ~ prefix, to look for ~/L/")
	logCmd.Flags().Bool("multiline", false, "turn on or off multiple line logging")
	logCmd.Flags().String("n", "", "number of initial lines to display")
	logCmd.Flags().Bool("page", false, "use paging")
	logCmd.Flags().String("path", "", "use <logs-dir> instead of default")
	logCmd.Flags().Bool("shorten", false, "shorten UUIDs, paths, etc")
	logCmd.Flags().Bool("wait", false, "wait for new logs continuously")

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})
}