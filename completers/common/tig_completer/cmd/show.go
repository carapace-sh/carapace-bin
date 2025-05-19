package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:                "show",
	Short:              "Show various types of objects",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(showCmd).Standalone()

	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = append([]string{"show"}, c.Args...)
			return bridge.ActionCarapaceBin("git").Invoke(c).ToA()
		}),
	)
}
