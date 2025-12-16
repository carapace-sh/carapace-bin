package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "numfmt",
	Short: "Convert numbers from/to human-readable strings",
	Long:  "https://www.man7.org/linux/man-pages/man1/numfmt.1.html",
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

	rootCmd.Flags().Bool("debug", false, "print warnings about invalid input")
	rootCmd.Flags().StringP("delimiter", "d", "", "use X instead of whitespace for field delimiter")
	rootCmd.Flags().String("field", "", "replace the numbers in these input fields (default=1)")
	rootCmd.Flags().String("format", "", "use printf style floating-point FORMAT")
	rootCmd.Flags().String("from", "", "auto-scale input numbers to UNITs")
	rootCmd.Flags().String("from-unit", "", "specify the input unit size (instead of the default 1)")
	rootCmd.Flags().Bool("grouping", false, "use locale-defined grouping of digits")
	rootCmd.Flags().String("header", "", "print (without converting) the first N header lines")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("invalid", "", "failure mode for invalid numbers")
	rootCmd.Flags().String("padding", "", "pad the output to N characters")
	rootCmd.Flags().String("round", "", "use METHOD for rounding when scaling")
	rootCmd.Flags().String("suffix", "", "add SUFFIX to output numbers, and accept optional SUFFIX in input numbers")
	rootCmd.Flags().String("to", "", "auto-scale output numbers to UNITs")
	rootCmd.Flags().String("to-unit", "", "the output unit size (instead of the default 1)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	rootCmd.Flag("header").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"from":    carapace.ActionValues("none", "auto", "si", "iec", "iec-i"),
		"invalid": carapace.ActionValues("abort", "fail", "warn", "ignore"),
		"round":   carapace.ActionValues("up", "down", "from-zero", "towards-zero", "nearest"),
		"to":      carapace.ActionValues("none", "auto", "si", "iec", "iec-i"),
	})
}
