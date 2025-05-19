package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [OPTIONS] [SERVICE...]",
	Short: "Removes stopped service containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().BoolP("all", "a", false, "Deprecated - no effect")
	rmCmd.Flags().BoolP("force", "f", false, "Don't ask to confirm removal")
	rmCmd.Flags().BoolP("stop", "s", false, "Stop the containers, if required, before removing")
	rmCmd.Flags().BoolP("volumes", "v", false, "Remove any anonymous volumes attached to containers")
	rmCmd.Flag("all").Hidden = true
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		action.ActionServices(rmCmd).FilterArgs(),
	)
}
