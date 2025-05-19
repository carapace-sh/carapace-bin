package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var azCmd = &cobra.Command{
	Use:   "az",
	Short: "Access Azure API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(azCmd).Standalone()

	azCmd.Flags().String("app", "", "Optional name of the Azure application to use if logged into multiple.")
	rootCmd.AddCommand(azCmd)

	azCmd.Flags().SetInterspersed(false)

	// TODO proxy the az command
	carapace.Gen(azCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("az"),
	)
}
