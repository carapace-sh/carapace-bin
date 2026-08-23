package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinedirectoryCmd = &cobra.Command{
	Use:   "machinedirectory",
	Short: "print the path to the current machine directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinedirectoryCmd).Standalone()
	rootCmd.AddCommand(machinedirectoryCmd)
}