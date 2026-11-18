package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a thread and its messages. Accepts an id, a harness id or `latest`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_showCmd).Standalone()

	threads_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	threads_showCmd.Flags().String("style", "auto", "How output is rendered")
	threadsCmd.AddCommand(threads_showCmd)

	carapace.Gen(threads_showCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
