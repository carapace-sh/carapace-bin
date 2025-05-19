package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "OpenFaaS store commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storeCmd).Standalone()
	storeCmd.PersistentFlags().StringP("platform", "p", "", "Target platform for store")
	storeCmd.PersistentFlags().StringP("url", "u", "https://raw.githubusercontent.com/openfaas/store/master/functions.json", "Alternative Store URL starting with http(s)://")
	rootCmd.AddCommand(storeCmd)

	carapace.Gen(storeCmd).FlagCompletion(carapace.ActionMap{
		"platform": action.ActionImageArchitectures(), // TODO verify
	})
}
