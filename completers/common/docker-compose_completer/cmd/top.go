package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top [SERVICES...]",
	Short: "Display the running processes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(topCmd).Standalone()

	rootCmd.AddCommand(topCmd)

	carapace.Gen(topCmd).PositionalAnyCompletion(
		action.ActionServices(topCmd).FilterArgs(),
	)
}
