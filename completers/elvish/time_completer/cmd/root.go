package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "time",
	Short: "Time execution of a command",
	Long:  "https://elv.sh/ref/builtin.html#time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("on-end", "on-end", "", "callback with duration on end")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
