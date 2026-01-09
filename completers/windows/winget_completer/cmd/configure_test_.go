package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_testCmd = &cobra.Command{
	Use:   "test",
	Short: "Checks the system against a desired state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_testCmd).Standalone()

	configureCmd.AddCommand(configure_testCmd)
}
