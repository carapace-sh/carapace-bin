package cmd

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "montage",
	Short:              "create a composite image by combining several separate images",
	Long:               "https://imagemagick.org/script/montage.php",
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
				return bridge.ActionCarapace("carapace-magick", "montage")
			}
			return bridge.ActionBridge("montage")
		}),
	)
}
