package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tail",
	Short: "display the last part of a file",
	Long:  "https://keith.github.io/xcode-manpages/tail.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bytes", "c", "", "Print the last n bytes of each file")
	rootCmd.Flags().BoolP("follow", "f", false, "Follow the file as it grows")
	rootCmd.Flags().StringP("lines", "n", "", "Print the last n lines of each file")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print headers indicating file names")
	rootCmd.Flags().BoolP("reverse", "r", false, "Print the lines in reverse order")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print headers indicating file names")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
