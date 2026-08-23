package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pbcopy",
	Short: "provide copying to the pasteboard",
	Long:  "https://keith.github.io/xcode-manpages/pbcopy.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("pboard", "p", "", "Specify pasteboard: general, ruler, find, or font")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pboard": carapace.ActionValues("general", "ruler", "find", "font"),
	})
}
