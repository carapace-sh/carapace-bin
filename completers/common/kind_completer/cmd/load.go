package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads images into nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadCmd).Standalone()

	rootCmd.AddCommand(loadCmd)
}
