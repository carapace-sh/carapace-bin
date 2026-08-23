package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskinfo",
	Short: "display policy information from kernel",
	Long:  "https://keith.github.io/xcode-manpages/taskinfo.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("boosts", false, "Display boost information")
	rootCmd.Flags().Bool("dq", false, "Display dispatch queue information")
	rootCmd.Flags().Bool("threads", false, "Display thread information")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.Batch(
		carapace.ActionValues(),
		ps.ActionProcessIds(),
	).ToA())
}
