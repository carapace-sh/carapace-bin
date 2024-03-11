package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var gsutilCmd = &cobra.Command{
	Use:   "gsutil",
	Short: "Access Google Cloud Storage with the gsutil command.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gsutilCmd).Standalone()

	gsutilCmd.Flags().String("app", "", "Optional name of the GCP application to use if logged into multiple.")
	rootCmd.AddCommand(gsutilCmd)

	gsutilCmd.Flags().SetInterspersed(false)

	// TODO proxy the gsutil command
	carapace.Gen(gsutilCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gsutil"),
	)
}
