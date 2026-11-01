package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var windowCmd = &cobra.Command{
	Use:   "window",
	Short: "Manage windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowCmd).Standalone()

	rootCmd.AddCommand(windowCmd)
}
