package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [OPTIONS] <DOMAIN_NAME> [PROG]...",
	Short: "Connect to wezterm multiplexer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectCmd).Standalone()
	connectCmd.Flags().SetInterspersed(false)

	connectCmd.Flags().String("class", "", "Override the default windowing system class")
	connectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	connectCmd.Flags().String("position", "", "Override the position for the initial window launched by this process")
	connectCmd.Flags().String("workspace", "", "Override the default workspace with the provided name")
	rootCmd.AddCommand(connectCmd)

	carapace.Gen(connectCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO domain
		bridge.ActionCarapaceBin().Shift(1),
	)
}
