package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "renice",
	Short: "alter priority of running processes",
	Long:  "https://keith.github.io/xcode-manpages/renice.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("g", "g", false, "Interpret target parameters as process group ID's")
	rootCmd.Flags().BoolS("n", "n", false, "Interpret the following argument as an increment")
	rootCmd.Flags().BoolS("p", "p", false, "Interpret target parameters as process ID's")
	rootCmd.Flags().BoolS("u", "u", false, "Interpret target parameters as user names or user ID's")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.Batch(
		ps.ActionProcessIds(),
		os.ActionUsers(),
	).ToA())
}
