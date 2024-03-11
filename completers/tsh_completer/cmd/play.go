package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Replay the recorded session (SSH, Kubernetes, App, DB).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(playCmd).Standalone()

	playCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	playCmd.Flags().StringP("format", "f", "", "Format output (pty, json, yaml)")
	rootCmd.AddCommand(playCmd)

	carapace.Gen(playCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"format":  tsh.ActionFormats(),
	})
}
