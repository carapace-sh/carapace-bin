package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killAllCmd = &cobra.Command{
	Use:   "kill-all",
	Short: "Kill all processes that are safe to kill (cached, etc)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killAllCmd).Standalone()

	rootCmd.AddCommand(killAllCmd)

}
