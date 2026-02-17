package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dumpCmd).Standalone()

	helpCmd.AddCommand(help_dumpCmd)
}
