package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitStatusCmd = &cobra.Command{
	Use:   "exit-status",
	Short: "List exit status definitions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitStatusCmd).Standalone()

	rootCmd.AddCommand(exitStatusCmd)
}
