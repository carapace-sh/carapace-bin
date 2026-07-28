package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitNodeCmd = &cobra.Command{
	Use:   "exit-node",
	Short: "Show machines on your tailnet configured as exit nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitNodeCmd).Standalone()

	rootCmd.AddCommand(exitNodeCmd)
}
