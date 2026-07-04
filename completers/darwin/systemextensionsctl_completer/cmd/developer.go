package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var developerCmd = &cobra.Command{
	Use:   "developer",
	Short: "Enable or disable developer mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(developerCmd).Standalone()
	rootCmd.AddCommand(developerCmd)
	carapace.Gen(developerCmd).PositionalCompletion(
		carapace.ActionValues("on", "off"),
	)
}
