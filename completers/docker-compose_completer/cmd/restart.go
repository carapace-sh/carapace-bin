package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [OPTIONS] [SERVICE...]",
	Short: "Restart service containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()

	restartCmd.Flags().Bool("no-deps", false, "Don't restart dependent services.")
	restartCmd.Flags().IntP("timeout", "t", 10, "Specify a shutdown timeout in seconds")
	rootCmd.AddCommand(restartCmd)

	carapace.Gen(restartCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(restartCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
