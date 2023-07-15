package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "run specified go tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolCmd).Standalone()

	toolCmd.Flags().BoolS("n", "n", false, "only print the command that would be executed")

	toolCmd.Flags().SetInterspersed(false)
	rootCmd.AddCommand(toolCmd)

	carapace.Gen(toolCmd).PositionalCompletion(
		golang.ActionTools(),
	)

	carapace.Gen(toolCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin("go-tool-" + c.Args[0]).Shift(1)
		}),
	)
}
