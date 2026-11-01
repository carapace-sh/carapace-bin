package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clCmd = &cobra.Command{
	Use:   "cl",
	Short: "clear an event log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clCmd).Standalone()
	rootCmd.AddCommand(clCmd)
}
