package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "manage directory partitions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directoryCmd).Standalone()
	rootCmd.AddCommand(directoryCmd)
}
