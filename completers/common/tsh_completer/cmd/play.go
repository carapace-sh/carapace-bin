package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Replay the recorded session (SSH, Kubernetes, App, DB).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(playCmd).Standalone()

	playCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	playCmd.Flags().StringP("format", "f", "pty", "Format output (pty, json, yaml, text).")
	playCmd.Flags().Bool("no-skip-idle-time", false, "Quickly skip over idle time, applicable when streaming SSH or Kubernetes sessions.")
	playCmd.Flags().Bool("skip-idle-time", false, "Quickly skip over idle time, applicable when streaming SSH or Kubernetes sessions.")
	playCmd.Flags().String("speed", "1x", "Playback speed, applicable when streaming SSH or Kubernetes sessions.")
	playCmd.Flag("no-skip-idle-time").Hidden = true
	rootCmd.AddCommand(playCmd)
}
