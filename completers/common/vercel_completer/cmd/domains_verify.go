package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var domains_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check a domain's DNS configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_verifyCmd).Standalone()

	domains_verifyCmd.Flags().String("format", "", "Output format")
	domains_verifyCmd.Flags().Bool("json", false, "Output as JSON")
	domains_verifyCmd.Flags().String("project", "", "Project name or ID")
	domains_verifyCmd.Flags().Bool("strict", false, "Strict mode")

	domainsCmd.AddCommand(domains_verifyCmd)

	carapace.Gen(domains_verifyCmd).FlagCompletion(carapace.ActionMap{
		"format":  carapace.ActionValues("plain", "json"),
		"project": action.ActionProjects(domains_verifyCmd),
	})

	carapace.Gen(domains_verifyCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
