package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [SERVICE...]",
	Short: "Watch build context for service and rebuild/refresh containers when files are updated",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().Bool("no-up", false, "Do not build & start services before watching")
	watchCmd.Flags().Bool("prune", false, "Prune dangling images on rebuild")
	watchCmd.Flags().Bool("quiet", false, "hide build output")
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).PositionalAnyCompletion(
		action.ActionServices(watchCmd).FilterArgs(),
	)
}
