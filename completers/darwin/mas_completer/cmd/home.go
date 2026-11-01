package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Open app page in the default web browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()
	homeCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	rootCmd.AddCommand(homeCmd)
}
