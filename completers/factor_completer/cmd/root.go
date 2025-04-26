package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "factor",
	Short: "factor numbers",
	Long:  "https://man7.org/linux/man-pages/man1/factor.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("exponents", "h", false, "print repeated factors in form p^e unless e is 1")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
