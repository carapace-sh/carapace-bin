package cmd

import (
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "run specified go tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolCmd).Standalone()
	toolCmd.Flags().SetInterspersed(false)

	toolCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	toolCmd.Flags().BoolS("n", "n", false, "only print the command that would be executed")
	rootCmd.AddCommand(toolCmd)

	carapace.Gen(toolCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})

	carapace.Gen(toolCmd).PositionalCompletion(
		golang.ActionTools(),
	)

	carapace.Gen(toolCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "pprof":
				return bridge.ActionCarapaceBin(c.Args[0]).Shift(1)
			default:
				return bridge.ActionCarapaceBin("go-tool-" + filepath.Base(c.Args[0])).Shift(1)
			}
		}),
	)

	carapace.Gen(toolCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("C").Value.String())
	})
}
