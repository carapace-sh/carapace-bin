package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off logging for all prompts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_offCmd).Standalone()

	logsCmd.AddCommand(logs_offCmd)
}
