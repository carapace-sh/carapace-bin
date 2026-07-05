package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gliCmd = &cobra.Command{
	Use:   "gli",
	Short: "get event log status information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gliCmd).Standalone()
	rootCmd.AddCommand(gliCmd)
}
