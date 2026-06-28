package cmd

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "compare",
	Short:              "mathematically and visually annotate the difference between an image and its reconstruction",
	Long:               "https://imagemagick.org/script/compare.php",
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
				return bridge.ActionCarapace("carapace-magick", "compare")
			}
			return bridge.ActionBridge("compare")
		}),
	)
}
