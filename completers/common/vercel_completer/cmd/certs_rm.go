package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certs_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove a certificate by id",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_rmCmd).Standalone()

	certsCmd.AddCommand(certs_rmCmd)

	carapace.Gen(certs_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
