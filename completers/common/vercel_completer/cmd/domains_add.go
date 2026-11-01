package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var domains_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a domain name that you already own",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_addCmd).Standalone()

	domains_addCmd.Flags().Bool("force", false, "Force add the domain")

	domainsCmd.AddCommand(domains_addCmd)

	carapace.Gen(domains_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionProjects(domains_addCmd),
	)
}
