package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Send data to one or more plugins, launch them if they are not running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pipeCmd).Standalone()

	helpCmd.AddCommand(help_pipeCmd)
}
