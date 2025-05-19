package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "Operate on runners",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerCmd).Standalone()
	rootCmd.AddCommand(runnerCmd)
}
