package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn on logging for all prompts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_onCmd).Standalone()

	logsCmd.AddCommand(logs_onCmd)
}
