package cmd

import (
	"github.com/rsteube/carapace"
	man "github.com/rsteube/carapace-bin/completers/man_completer/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "batman",
	Short:              "Read system manual pages (man) using bat",
	Long:               "https://github.com/eth-p/bat-extras/blob/master/doc/batman.md",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		man.ActionExecute(),
	)
}
