package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var killSessionCmd = &cobra.Command{
	Use:     "kill-session",
	Short:   "Kill a specific session",
	Aliases: []string{"k"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killSessionCmd).Standalone()

	killSessionCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(killSessionCmd)

	carapace.Gen(killSessionCmd).PositionalCompletion(
		zellij.ActionSessions(),
	)
}
