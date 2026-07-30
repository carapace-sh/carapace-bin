package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load a subkey into a different subkey",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadCmd).Standalone()
	rootCmd.AddCommand(loadCmd)
}
