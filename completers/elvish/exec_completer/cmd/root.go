package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "exec",
	Short: "Replace the current Elvish process with an external command",
	Long:  "https://elv.sh/ref/builtin.html#exec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionExecutables(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
