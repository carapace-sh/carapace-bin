package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons___providerExecCmd = &cobra.Command{
	Use:    "__provider-exec",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons___providerExecCmd).Standalone()

	daemons___providerExecCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons___providerExecCmd)
}
