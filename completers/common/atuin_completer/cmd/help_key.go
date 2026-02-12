package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Print the encryption key for transfer to another machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_keyCmd).Standalone()

	helpCmd.AddCommand(help_keyCmd)
}
