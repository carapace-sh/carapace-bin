package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "file specific operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fileCmd).Standalone()
	rootCmd.AddCommand(fileCmd)
}
