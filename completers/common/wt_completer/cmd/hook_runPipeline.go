package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_runPipelineCmd = &cobra.Command{
	Use:    "run-pipeline",
	Short:  "Internal: run a serialized pipeline from stdin",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_runPipelineCmd).Standalone()

	hook_runPipelineCmd.Flags().BoolP("help", "h", false, "Print help")
	hookCmd.AddCommand(hook_runPipelineCmd)
}
