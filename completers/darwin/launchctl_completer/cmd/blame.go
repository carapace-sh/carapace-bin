package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Print the reason a service is running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()
	rootCmd.AddCommand(blameCmd)
	carapace.Gen(blameCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
