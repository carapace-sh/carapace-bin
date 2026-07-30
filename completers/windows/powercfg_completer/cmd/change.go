package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "modify a setting in the current power plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changeCmd).Standalone()
	rootCmd.AddCommand(changeCmd)
}
