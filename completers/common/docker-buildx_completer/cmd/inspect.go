package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-buildx_completer/cmd/action"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [NAME]",
	Short: "Inspect current builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()
	inspectCmd.Flags().Bool("bootstrap", false, "Ensure builder has booted before inspecting")
	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).PositionalCompletion(
		action.ActionBuilders(),
	)
}
