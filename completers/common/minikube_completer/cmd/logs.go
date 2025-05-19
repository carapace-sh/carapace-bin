package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:     "logs",
	Short:   "Returns logs to debug a local Kubernetes cluster",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()
	logsCmd.Flags().String("file", "", "If present, writes to the provided file instead of stdout.")
	logsCmd.Flags().BoolP("follow", "f", false, "Show only the most recent journal entries, and continuously print new entries as they are appended to the journal.")
	logsCmd.Flags().IntP("length", "n", 60, "Number of lines back to go within the log")
	logsCmd.Flags().String("node", "", "The node to get logs from. Defaults to the primary control plane.")
	logsCmd.Flags().Bool("problems", false, "Show only log entries which point to known problems")
	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"node": action.ActionNodes(),
	})
}
