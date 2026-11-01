package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List every setting and shortcut with its current value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configListCmd).Standalone()

	configCmd.AddCommand(configListCmd)
}
