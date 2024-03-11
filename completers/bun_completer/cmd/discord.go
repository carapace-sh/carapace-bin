package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discordCmd = &cobra.Command{
	Use:   "discord",
	Short: "Open bun's Discord server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discordCmd).Standalone()

	rootCmd.AddCommand(discordCmd)
}
