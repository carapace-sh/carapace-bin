package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "date",
	Short: "display or set date and time",
	Long:  "https://keith.github.io/xcode-manpages/date.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("adjust", "v", "", "Adjust (i.e., take the current date and time and add or subtract the specified value)")
	rootCmd.Flags().StringP("format", "f", "", "Use fmt as the format string to parse the input")
	rootCmd.Flags().BoolP("j", "j", false, "Do not try to set the date")
	rootCmd.Flags().BoolP("n", "n", false, "Indicate that the time is specified in seconds since the Epoch")
	rootCmd.Flags().StringP("reference", "r", "", "Print the date and time of the file referenced by file")
	rootCmd.Flags().StringP("seconds", "s", "", "Set the date and time to the time specified by seconds")
	rootCmd.Flags().BoolP("u", "u", false, "Display or set the date in UTC (Coordinated Universal) time")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Prefix("+"),
	)
}
