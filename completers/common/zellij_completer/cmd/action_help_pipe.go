package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Send data to one or more plugins, launch them if they are not running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_pipeCmd).Standalone()

	action_helpCmd.AddCommand(action_help_pipeCmd)
}
