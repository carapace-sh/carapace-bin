package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a service (deprecated, use kickstart)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()
	rootCmd.AddCommand(startCmd)
	carapace.Gen(startCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
