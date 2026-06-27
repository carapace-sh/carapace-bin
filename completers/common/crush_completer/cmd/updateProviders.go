package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/crush"
	"github.com/spf13/cobra"
)

var updateProvidersCmd = &cobra.Command{
	Use:   "update-providers [path-or-url]",
	Short: "Update providers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateProvidersCmd).Standalone()

	updateProvidersCmd.Flags().String("source", "", "Provider source to update (catwalk or hyper)")
	rootCmd.AddCommand(updateProvidersCmd)

	carapace.Gen(updateProvidersCmd).FlagCompletion(carapace.ActionMap{
		"source": crush.ActionProviderSources(),
	})
	carapace.Gen(updateProvidersCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			http.ActionUrls(),
		).ToA(),
	)
}
