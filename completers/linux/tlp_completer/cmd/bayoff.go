package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bayoffCmd = &cobra.Command{
	Use:   "bayoff",
	Short: "Turn off optical drive in UltraBay/MediaBay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bayoffCmd).Standalone()

	rootCmd.AddCommand(bayoffCmd)

	carapace.Gen(bayoffCmd).DashAnyCompletion(action.ActionConfigParams())
}
