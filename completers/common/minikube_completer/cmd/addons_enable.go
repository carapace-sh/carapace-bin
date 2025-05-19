package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var addons_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables the addon w/ADDON_NAME within minikube. For a list of available addons use: minikube addons list ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addons_enableCmd).Standalone()
	addons_enableCmd.Flags().Bool("force", false, "If true, will perform potentially dangerous operations. Use with discretion.")
	addons_enableCmd.Flags().String("images", "", "Images used by this addon. Separated by commas.")
	addons_enableCmd.Flags().Bool("refresh", false, "If true, pods might get deleted and restarted on addon enable")
	addons_enableCmd.Flags().String("registries", "", "Registries used by this addon. Separated by commas.")
	addonsCmd.AddCommand(addons_enableCmd)

	carapace.Gen(addons_enableCmd).FlagCompletion(carapace.ActionMap{
		"images": docker.ActionRepositoryTags().UniqueList(","),
	})

	carapace.Gen(addons_enableCmd).PositionalCompletion(
		action.ActionAddons(),
	)
}
