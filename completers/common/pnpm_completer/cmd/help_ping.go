package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connectivity to the configured registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pingCmd).Standalone()

	helpCmd.AddCommand(help_pingCmd)
}
