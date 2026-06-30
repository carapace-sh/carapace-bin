package cmd

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "ffmpeg",
	Short:              "Hyper fast Audio and Video encoder",
	Long:               "https://ffmpeg.org/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if _, err := exec.LookPath("carapace-ffmpeg"); err == nil {
				return bridge.ActionCarapace("carapace-ffmpeg", "ffmpeg")
			}
			return bridge.ActionBridge("ffmpeg")
		}),
	)
}
