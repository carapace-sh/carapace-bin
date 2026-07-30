package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "synchronize a directory to another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()
	rootCmd.AddCommand(syncCmd)
}
