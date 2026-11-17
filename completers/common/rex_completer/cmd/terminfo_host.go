package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Install Rex's terminfo entry on other machines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoHostCmd).Standalone()

	terminfoCmd.AddCommand(terminfoHostCmd)
}
