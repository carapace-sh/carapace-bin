package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var gcloudCmd = &cobra.Command{
	Use:     "gcloud",
	Short:   "Access GCP API with the gcloud command.",
	Aliases: []string{"gcp"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcloudCmd).Standalone()

	gcloudCmd.Flags().String("app", "", "Optional name of the GCP application to use if logged into multiple.")
	rootCmd.AddCommand(gcloudCmd)

	gcloudCmd.Flags().SetInterspersed(false)

	// TODO proxy the gcloud command
	carapace.Gen(gcloudCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gcloud"),
	)
}
