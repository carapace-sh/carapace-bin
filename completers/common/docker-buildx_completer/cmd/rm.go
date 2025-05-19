package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-buildx_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [NAME]",
	Short: "Remove a builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()
	rmCmd.Flags().Bool("all-inactive", false, "Remove all inactive builders")
	rmCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	rmCmd.Flags().Bool("keep-daemon", false, "Keep the buildkitd daemon running")
	rmCmd.Flags().Bool("keep-state", false, "Keep BuildKit state")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(
		action.ActionBuilders(),
	)
}
