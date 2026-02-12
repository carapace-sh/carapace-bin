package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kv_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a saved value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_help_getCmd).Standalone()

	kv_helpCmd.AddCommand(kv_help_getCmd)
}
