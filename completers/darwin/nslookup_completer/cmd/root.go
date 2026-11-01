package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nslookup",
	Short: "query Internet name servers interactively",
	Long:  "https://keith.github.io/xcode-manpages/nslookup.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
