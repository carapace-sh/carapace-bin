package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:                "kitty",
	Short:              "The fast, feature rich terminal emulator",
	Long:               "https://sw.kovidgoyal.net/kitty/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionKitten("kitty"),
	)

	fs := pflag.NewFlagSet("directory", pflag.ContinueOnError)
	fs.StringP("directory", "s", "", "")
	fs.String("working-directory", "", "")
	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		fs.Parse(args)
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		if f := fs.Lookup("directory"); f.Changed {
			return action.Chdir(f.Value.String())
		}
		if f := fs.Lookup("working-directory"); f.Changed {
			return action.Chdir(f.Value.String())
		}
		return action
	})
}
