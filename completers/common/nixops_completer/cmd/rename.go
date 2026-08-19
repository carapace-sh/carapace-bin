package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename a machine in the network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameCmd).Standalone()
	rootCmd.AddCommand(renameCmd)
}
