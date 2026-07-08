package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-effort",
	Short: "Displays effort statistics (number of commits per file)",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-effort",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("above", "", "Ignore files with commits below or equal to value")
	rootCmd.Flags().Bool("help", false, "show help")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
