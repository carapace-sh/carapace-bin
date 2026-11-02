package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keys_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Save a key in the keys.json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keys_setCmd).Standalone()

	keys_setCmd.Flags().String("value", "", "Value to set")
	keysCmd.AddCommand(keys_setCmd)
}
