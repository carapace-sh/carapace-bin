package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheClearCmd = &cobra.Command{
	Use:   "cache:clear [--tags [TAGS]] [--] [<store>]",
	Short: "Flush the application cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheClearCmd).Standalone()

	cacheClearCmd.Flags().String("tags", "", "The cache tags you would like to clear")
	rootCmd.AddCommand(cacheClearCmd)
}
