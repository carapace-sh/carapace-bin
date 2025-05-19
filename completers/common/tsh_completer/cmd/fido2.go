package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fido2Cmd = &cobra.Command{
	Use:    "fido2",
	Short:  "FIDO2 commands",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fido2Cmd).Standalone()

	rootCmd.AddCommand(fido2Cmd)
}
