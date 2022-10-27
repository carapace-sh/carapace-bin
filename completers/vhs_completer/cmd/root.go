package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vhs",
	Short: "Run a given tape file and generates its outputs.",
	Long:  "https://github.com/charmbracelet/vhs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "help for vhs")
	rootCmd.Flags().BoolP("version", "v", false, "version for vhs")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".tape"),
	)
}
