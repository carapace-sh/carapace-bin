package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var specs_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "display help for command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(specs_helpCmd).Standalone()

	specsCmd.AddCommand(specs_helpCmd)
}
