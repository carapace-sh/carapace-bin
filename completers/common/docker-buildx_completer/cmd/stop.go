package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-buildx_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [NAME]",
	Short: "Stop builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()
	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).PositionalCompletion(
		action.ActionBuilders(),
	)
}
