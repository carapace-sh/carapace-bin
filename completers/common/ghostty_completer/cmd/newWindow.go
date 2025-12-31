package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var newWindowCmd = &cobra.Command{
	Use:   "+new-window",
	Short: "The `new-window` will use native platform IPC to open up a new window in a running instance of Ghostty",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newWindowCmd).Standalone()
	newWindowCmd.Flags().SetInterspersed(false)

	newWindowCmd.Flags().String("class", "", "If set, open up a new window in a custom instance of Ghostty")
	newWindowCmd.Flags().BoolS("e", "e", false, "Any arguments after this will be interpreted as a command to execute")
	rootCmd.AddCommand(newWindowCmd)

	carapace.Gen(newWindowCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if newWindowCmd.Flag("e").Changed {
				return bridge.ActionCarapaceBin()
			}
			return carapace.ActionValues()
		}),
	)
}
