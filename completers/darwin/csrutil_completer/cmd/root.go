package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csrutil",
	Short: "configure system security policies",
	Long:  "https://keith.github.io/xcode-manpages/csrutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.Batch(
		carapace.ActionValues("clear", "enable", "disable"),
		carapace.ActionValuesDescribed(
			"status", "check SIP status",
			"netboot", "configure netboot settings",
		),
	).ToA())
}
