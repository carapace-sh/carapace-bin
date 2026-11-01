package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domains_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove ownership of a domain name",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_rmCmd).Standalone()

	domains_rmCmd.Flags().Bool("yes", false, "Skip confirmation")

	domainsCmd.AddCommand(domains_rmCmd)

	carapace.Gen(domains_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
