package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [OPTIONS] [SERVICE...]",
	Short: "Restart service containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()

	restartCmd.Flags().Bool("no-deps", false, "Don't restart dependent services")
	restartCmd.Flags().StringP("timeout", "t", "", "Specify a shutdown timeout in seconds")
	rootCmd.AddCommand(restartCmd)

	carapace.Gen(restartCmd).PositionalAnyCompletion(
		action.ActionServices(restartCmd).FilterArgs(),
	)
}
