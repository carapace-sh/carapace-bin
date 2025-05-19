package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait SERVICE [SERVICE...] [OPTIONS]",
	Short: "Block until containers of all (or specified) services stop.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().Bool("down-project", false, "Drops project when the first container stops")
	rootCmd.AddCommand(waitCmd)

	carapace.Gen(waitCmd).PositionalAnyCompletion(
		action.ActionServices(waitCmd).FilterArgs(),
	)
}
