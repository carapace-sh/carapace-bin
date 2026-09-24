package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:     "start-server",
	Aliases: []string{"start"},
	Short:   "start a tmux server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startServerCmd).Standalone()

	rootCmd.AddCommand(startServerCmd)
}
