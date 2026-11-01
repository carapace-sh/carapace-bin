package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "release the IP address for an adapter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()
	rootCmd.AddCommand(releaseCmd)
}
