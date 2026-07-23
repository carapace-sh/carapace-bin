package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-sed",
	Short: "Run grep and replace patterns in files",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-sed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("commit", "c", false, "Commit changes")
	rootCmd.Flags().StringP("flags", "f", "", "sed regex flags")
	rootCmd.Flags().BoolP("help", "h", false, "show help")

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
	)
}
