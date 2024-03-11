package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "Help about the command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)
}
