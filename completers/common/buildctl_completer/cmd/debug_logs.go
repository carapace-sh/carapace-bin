package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "display build logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_logsCmd).Standalone()

	debug_logsCmd.Flags().String("progress", "", "progress output type")
	debug_logsCmd.Flags().Bool("trace", false, "show opentelemetry trace")
	debugCmd.AddCommand(debug_logsCmd)

	carapace.Gen(debug_logsCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto"),
	})
}
