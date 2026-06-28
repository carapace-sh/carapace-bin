package cmd

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "mogrify",
	Short:              "resize an image, blur, crop, despeckle, dither, draw on, flip, join, re-sample, and much more",
	Long:               "https://imagemagick.org/script/mogrify.php",
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
			if _, err := exec.LookPath("carapace-magick"); err == nil {
				return bridge.ActionCarapace("carapace-magick", "mogrify")
			}
			return bridge.ActionBridge("mogrify")
		}),
	)
}
