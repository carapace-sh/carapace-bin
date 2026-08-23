package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var foreverCmd = &cobra.Command{
	Use:   "forever",
	Short: "Wait forever before executing the next command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(foreverCmd).Standalone()
	rootCmd.AddCommand(foreverCmd)
}
