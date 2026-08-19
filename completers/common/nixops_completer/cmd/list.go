package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ListCmd).Standalone()
	rootCmd.AddCommand(ListCmd)
}
