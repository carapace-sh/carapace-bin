package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Manages custom dependency policies and their sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policyCmd).Standalone()

	policyCmd.Flags().StringP("file", "f", "", "If you want to choose a different composer.json or config.json")
	policyCmd.Flags().BoolP("global", "g", false, "Apply command to the global config file")

	rootCmd.AddCommand(policyCmd)

	carapace.Gen(policyCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(".json"),
	})

	carapace.Gen(policyCmd).PositionalCompletion(
		carapace.ActionValues("add-source"),
		action.ActionPolicyNames(policyCmd),
		action.ActionPolicySourceTypes(),
	)
}
