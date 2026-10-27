package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Rename or move tracked paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_mvCmd).Standalone()

	helpCmd.AddCommand(help_mvCmd)
}
