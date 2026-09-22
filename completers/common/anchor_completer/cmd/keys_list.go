package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keys_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the program keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keys_listCmd).Standalone()

	keys_listCmd.Flags().BoolP("help", "h", false, "Print help")
	keysCmd.AddCommand(keys_listCmd)
}
