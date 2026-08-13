package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var domains_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move ownership of a domain name to another Vercel Team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_moveCmd).Standalone()

	domains_moveCmd.Flags().Bool("yes", false, "Skip confirmation")

	domainsCmd.AddCommand(domains_moveCmd)

	carapace.Gen(domains_moveCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionTeams(domains_moveCmd),
	)
}
