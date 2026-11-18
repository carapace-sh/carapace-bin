package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List captured threads, newest first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_listCmd).Standalone()

	threads_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	threads_listCmd.Flags().String("style", "auto", "How output is rendered")
	threadsCmd.AddCommand(threads_listCmd)

	carapace.Gen(threads_listCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
