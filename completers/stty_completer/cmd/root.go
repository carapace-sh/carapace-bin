package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stty",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Root()
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print all current settings in human-readable form")
	rootCmd.Flags().StringP("file", "F", "", "open and use the specified DEVICE instead of stdin")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("save", "g", false, "print all current settings in a stty-readable form")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
