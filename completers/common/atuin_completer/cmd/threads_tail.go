package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Follow threads and messages as they are recorded (until interrupted)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_tailCmd).Standalone()

	threads_tailCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	threads_tailCmd.Flags().String("style", "auto", "How output is rendered")
	threadsCmd.AddCommand(threads_tailCmd)

	carapace.Gen(threads_tailCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
