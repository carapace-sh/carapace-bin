package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:    "all",
	Short:  "List all commands, including hidden ones",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allCmd).Standalone()

	rootCmd.AddCommand(allCmd)
}
