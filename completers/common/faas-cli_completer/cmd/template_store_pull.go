package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var template_store_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull templates from store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(template_store_pullCmd).Standalone()
	template_store_pullCmd.PersistentFlags().StringP("url", "u", "https://raw.githubusercontent.com/openfaas/store/master/templates.json", "Use as alternative store for templates")
	template_storeCmd.AddCommand(template_store_pullCmd)

	carapace.Gen(template_store_pullCmd).PositionalCompletion(
		action.ActionLanguageTemplates(),
	)
}
