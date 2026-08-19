package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(StopCmd).Standalone()
	rootCmd.AddCommand(StopCmd)
}
