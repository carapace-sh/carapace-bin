package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_syncCmd).Standalone()

	helpCmd.AddCommand(help_syncCmd)
}
