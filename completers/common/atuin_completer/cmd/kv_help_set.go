package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kv_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a key-value pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_help_setCmd).Standalone()

	kv_helpCmd.AddCommand(kv_help_setCmd)
}
