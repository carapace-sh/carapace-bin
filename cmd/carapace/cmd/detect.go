package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var detectCmd = &cobra.Command{
	Use:   "--detect command",
	Short: "detect bridge by invoking command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if bridge, ok := bridge.Detect(args[0]); ok {
			cmd.Println(bridge.Name)
		} else {
			cmd.PrintErrln("none detected")
			os.Exit(1)
		}
	},
}

func init() {
	carapace.Gen(detectCmd).Standalone()

	carapace.Gen(detectCmd).PositionalCompletion(
		carapace.ActionExecutables(),
	)
}
