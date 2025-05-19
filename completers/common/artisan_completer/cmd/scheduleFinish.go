package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleFinishCmd = &cobra.Command{
	Use:   "schedule:finish <id> [<code>]",
	Short: "Handle the completion of a scheduled command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleFinishCmd).Standalone()

	rootCmd.AddCommand(scheduleFinishCmd)
}
