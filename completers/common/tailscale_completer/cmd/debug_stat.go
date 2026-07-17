package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Stat a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_statCmd).Standalone()

	debugCmd.AddCommand(debug_statCmd)
}
