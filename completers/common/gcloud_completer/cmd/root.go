package cmd

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "gcloud",
	Short:              "manage Google Cloud Platform resources and developer workflow",
	Long:               "https://cloud.google.com/sdk/gcloud/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if _, err := exec.LookPath("carapace-gcloud"); err == nil {
				return bridge.ActionCarapace("carapace-gcloud")
			}
			return bridge.ActionGcloud("gcloud")
		}),
	)
}
