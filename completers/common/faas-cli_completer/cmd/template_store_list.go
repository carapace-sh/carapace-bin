package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var template_store_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List templates from OpenFaaS organizations",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(template_store_listCmd).Standalone()
	template_store_listCmd.Flags().BoolP("json", "j", false, "Output templates as JSON")
	template_store_listCmd.Flags().BoolP("official", "o", false, "Shows only official templates")
	template_store_listCmd.Flags().StringP("platform", "p", "x86_64", "Shows the platform if the output is verbose")
	template_store_listCmd.PersistentFlags().StringP("url", "u", "https://raw.githubusercontent.com/openfaas/store/master/templates.json", "Use as alternative store for templates")
	template_store_listCmd.Flags().BoolP("recommended", "r", false, "Shows only recommended templates")
	template_store_listCmd.Flags().BoolP("verbose", "v", false, "Shows additional language and platform")
	template_storeCmd.AddCommand(template_store_listCmd)

	carapace.Gen(template_store_listCmd).FlagCompletion(carapace.ActionMap{
		"platform": action.ActionArchitectures(),
	})
}
