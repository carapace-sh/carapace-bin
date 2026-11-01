package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
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
	gcloudCmd.Flags().String("gcp-service-account", "", "(For GCP CLI access only) GCP service account name.")
	rootCmd.AddCommand(gcloudCmd)

	gcloudCmd.Flags().SetInterspersed(false)

	carapace.Gen(gcloudCmd).FlagCompletion(carapace.ActionMap{
		"app": tsh.ActionApps(),
	})

	carapace.Gen(gcloudCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gcloud"),
	)
}
