package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate random numbers",
	Long:  "https://fishshell.com/docs/current/cmds/random.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("choice").Usage("select from a list of items"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 && c.Args[0] == "choice" {
				return bridge.ActionCarapaceBin()
			}
			return carapace.ActionValues()
		}),
	)
}
