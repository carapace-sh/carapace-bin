package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var procinfoCmd = &cobra.Command{
	Use:   "procinfo",
	Short: "Prints port information about a process",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(procinfoCmd).Standalone()
	rootCmd.AddCommand(procinfoCmd)
}
