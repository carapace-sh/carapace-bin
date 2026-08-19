package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RenameCmd).Standalone()
	rootCmd.AddCommand(RenameCmd)
}
