package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "*Experimental* Start the background daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_daemonCmd).Standalone()

	helpCmd.AddCommand(help_daemonCmd)
}
