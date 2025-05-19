package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var touchidCmd = &cobra.Command{
	Use:    "touchid",
	Short:  "Manage Touch ID credentials",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(touchidCmd).Standalone()

	rootCmd.AddCommand(touchidCmd)
}
