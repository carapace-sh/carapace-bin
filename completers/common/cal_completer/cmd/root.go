package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cal",
	Short: "display a calendar",
	Long:  "https://man7.org/linux/man-pages/man1/cal.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("color", "", "colorize messages")
	rootCmd.Flags().StringP("columns", "c", "", "amount of columns to use")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().Bool("iso", false, "alias for --reform=iso")
	rootCmd.Flags().BoolP("julian", "j", false, "use day-of-year for all calendars")
	rootCmd.Flags().BoolP("monday", "m", false, "Monday as first day of week")
	rootCmd.Flags().StringP("months", "n", "", "show num months starting with date's month")
	rootCmd.Flags().BoolP("one", "1", false, "show only a single month (default)")
	rootCmd.Flags().String("reform", "", "Gregorian reform date")
	rootCmd.Flags().BoolP("span", "S", false, "span the date when displaying multiple months")
	rootCmd.Flags().BoolP("sunday", "s", false, "Sunday as first day of week")
	rootCmd.Flags().BoolP("three", "3", false, "show three months spanning the date")
	rootCmd.Flags().BoolP("twelve", "Y", false, "show the next twelve months")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().BoolP("vertical", "v", false, "show day vertically instead of line")
	rootCmd.Flags().StringP("week", "w", "", "show US or ISO-8601 week numbers")
	rootCmd.Flags().BoolP("year", "y", false, "show the whole year")

	rootCmd.Flag("color").NoOptDefVal = " "
	rootCmd.Flag("week").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":  carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"reform": carapace.ActionValues("1752", "gregorian", "iso", "julian"),
	})
}
