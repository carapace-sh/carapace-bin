package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var indexSizeCmd = &cobra.Command{
	Use:   "indexSize",
	Short: "Display the disk space used by index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(indexSizeCmd).Standalone()

	rootCmd.AddCommand(indexSizeCmd)
}
