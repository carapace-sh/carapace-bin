package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limitCmd = &cobra.Command{
	Use:   "limit",
	Short: "Reads or modifies launchd's resource limits",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(limitCmd).Standalone()
	rootCmd.AddCommand(limitCmd)
}
