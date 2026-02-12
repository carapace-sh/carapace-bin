package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_kv_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a saved value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_kv_getCmd).Standalone()

	help_kvCmd.AddCommand(help_kv_getCmd)
}
