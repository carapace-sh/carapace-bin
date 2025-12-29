package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var archCmd = &cobra.Command{
	Use:   "arch",
	Short: "Prints the configured XBPS architecture",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(archCmd).Standalone()

	rootCmd.AddCommand(archCmd)
}
