package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_kv_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a key-value pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_kv_setCmd).Standalone()

	help_kvCmd.AddCommand(help_kv_setCmd)
}
