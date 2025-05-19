package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "patool",
	Short: "An archive file manager",
	Long:  "http://wummel.github.io/patool/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("non-interactive", false, "don't query for user input")
	rootCmd.Flags().CountP("quiet", "q", "quiet operation")
	rootCmd.Flags().CountP("verbose", "v", "verbose operation")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "show this help message and exit")
}
