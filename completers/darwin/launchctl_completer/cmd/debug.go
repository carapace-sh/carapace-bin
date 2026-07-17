package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Configure the next invocation of a service for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()
	rootCmd.AddCommand(debugCmd)
	carapace.Gen(debugCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
