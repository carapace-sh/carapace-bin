package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "unload a subkey",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unloadCmd).Standalone()
	rootCmd.AddCommand(unloadCmd)
}
