package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "head",
	Short: "display first lines of a file",
	Long:  "https://keith.github.io/xcode-manpages/head.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bytes", "c", "", "Print the first n bytes of each file")
	rootCmd.Flags().StringP("lines", "n", "", "Print the first n lines of each file")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print headers indicating file names")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print headers indicating file names")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
