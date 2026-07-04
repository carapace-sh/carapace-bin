package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installedCmd = &cobra.Command{
	Use:   "installed",
	Short: "List installed ports and their activation status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installedCmd).Standalone()
	rootCmd.AddCommand(installedCmd)
}
