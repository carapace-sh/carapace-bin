package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var keychainCmd = &cobra.Command{
	Use:   "keychain",
	Short: "Manipulate a device's keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keychainCmd).Standalone()
	rootCmd.AddCommand(keychainCmd)
	carapace.Gen(keychainCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("add-root-cert", "add-cert", "reset"),
	)
	carapace.Gen(keychainCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
