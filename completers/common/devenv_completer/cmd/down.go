package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop processes running in the background",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downCmd).Standalone()

	rootCmd.AddCommand(downCmd)
}
