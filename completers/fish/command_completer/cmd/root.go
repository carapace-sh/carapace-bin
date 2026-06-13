package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "command",
	Short: "Run a program ignoring functions",
	Long:  "https://fishshell.com/docs/current/cmds/command.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("a", "a", false, "print all matching commands")
	rootCmd.Flags().Bool("all", false, "print all matching commands")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("q", "q", false, "return 0 if any command found")
	rootCmd.Flags().Bool("query", false, "return 0 if any command found")
	rootCmd.Flags().BoolS("s", "s", false, "print path to external command")
	rootCmd.Flags().Bool("search", false, "print path to external command")
	rootCmd.Flags().BoolS("v", "v", false, "print path to external command")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionExecutables(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
