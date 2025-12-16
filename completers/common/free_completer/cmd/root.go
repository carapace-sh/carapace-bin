package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "free",
	Short: "Display amount of free and used memory in the system",
	Long:  "https://man7.org/linux/man-pages/man1/free.1.html",
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

	rootCmd.Flags().BoolP("bytes", "b", false, "show output in bytes")
	rootCmd.Flags().BoolP("committed", "v", false, "show committed memory and commit limit")
	rootCmd.Flags().StringP("count", "c", "", "repeat printing N times, then exit")
	rootCmd.Flags().BoolP("gibi", "g", false, "show output in gibibytes")
	rootCmd.Flags().Bool("giga", false, "show output in gigabytes")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("human", "h", false, "show human-readable output")
	rootCmd.Flags().BoolP("kibi", "k", false, "show output in kibibytes")
	rootCmd.Flags().Bool("kilo", false, "show output in kilobytes")
	rootCmd.Flags().BoolP("line", "L", false, "show output on a single line")
	rootCmd.Flags().BoolP("lohi", "l", false, "show detailed low and high memory statistics")
	rootCmd.Flags().BoolP("mebi", "m", false, "show output in mebibytes")
	rootCmd.Flags().Bool("mega", false, "show output in megabytes")
	rootCmd.Flags().Bool("pebi", false, "show output in pebibytes")
	rootCmd.Flags().Bool("peta", false, "show output in petabytes")
	rootCmd.Flags().StringP("seconds", "s", "", "repeat printing every N seconds")
	rootCmd.Flags().Bool("si", false, "use powers of 1000 not 1024")
	rootCmd.Flags().Bool("tebi", false, "show output in tebibytes")
	rootCmd.Flags().Bool("tera", false, "show output in terabytes")
	rootCmd.Flags().BoolP("total", "t", false, "show total for RAM + swap")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
	rootCmd.Flags().BoolP("wide", "w", false, "wide output")
}
