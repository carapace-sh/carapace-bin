package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename CONTAINER NAME",
	Short: "Rename an existing container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameCmd).Standalone()

	rootCmd.AddCommand(renameCmd)
}
