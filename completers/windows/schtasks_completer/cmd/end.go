package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "stop a running scheduled task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endCmd).Standalone()
	rootCmd.AddCommand(endCmd)
}
