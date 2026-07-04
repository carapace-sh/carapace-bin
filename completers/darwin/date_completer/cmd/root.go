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

	rootCmd.Flags().StringP("adjust", "v", "", "Adjust (add or subtract) the specified value from the current date")
	rootCmd.Flags().StringP("format", "f", "", "Use fmt as the format string to parse the input (with -j)")
	rootCmd.Flags().StringP("iso-8601", "I", "", "Output date/time in ISO 8601 format")
	rootCmd.Flags().BoolP("j", "j", false, "Do not try to set the date")
	rootCmd.Flags().BoolP("no-adjust", "n", false, "Do not adjust the time if the clock is out of sync")
	rootCmd.Flags().StringP("reference", "r", "", "Print the date and time of the file referenced by file")
	rootCmd.Flags().BoolP("rfc-2822", "R", false, "Output date and time in RFC 2822 format")
	rootCmd.Flags().BoolP("utc", "u", false, "Display or set the date in UTC (Coordinated Universal) time")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Prefix("+"),
	)
}
