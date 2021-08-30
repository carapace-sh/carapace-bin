package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/argcomplete"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:                "gcloud",
	Short:              "manage Google Cloud Platform resources and developer workflow",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO patch user@instance and --flag=optarg as in gcloud completion script

			if c.CallbackValue == "-" {
				return carapace.ActionValues("--").NoSpace() // seems shorthand flags aren't completed anyway so expand to longhand first
			}
			os.Setenv("CLOUDSDK_COMPONENT_MANAGER_DISABLE_UPDATE_CHECK", "1")
			return argcomplete.ActionArgcomplete("gcloud")
		}),
	)
}
