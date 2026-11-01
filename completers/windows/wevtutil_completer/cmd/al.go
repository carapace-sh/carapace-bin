package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alCmd = &cobra.Command{
	Use:   "al",
	Short: "archive a log file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alCmd).Standalone()
	rootCmd.AddCommand(alCmd)
}
