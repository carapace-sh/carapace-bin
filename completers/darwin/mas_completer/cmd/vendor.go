package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vendorCmd = &cobra.Command{
	Use:     "vendor",
	Aliases: []string{"seller"},
	Short:   "Open app seller pages in the default web browser",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vendorCmd).Standalone()
	vendorCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	rootCmd.AddCommand(vendorCmd)
}
