package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flushdnsCmd = &cobra.Command{
	Use:   "flushdns",
	Short: "purge the DNS resolver cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flushdnsCmd).Standalone()
	rootCmd.AddCommand(flushdnsCmd)
}
