package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "showkey",
	Short: "examine the codes sent by the keyboard",
	Long:  "https://linux.die.net/man/1/showkey",
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

	rootCmd.Flags().BoolP("ascii", "a", false, "display the decimal/octal/hex values of the keys")
	rootCmd.Flags().BoolP("help", "h", false, "print this usage message")
	rootCmd.Flags().BoolP("keycodes", "k", false, "display only the interpreted keycodes (default)")
	rootCmd.Flags().BoolP("scancodes", "s", false, "display only the raw scan-codes")
	rootCmd.Flags().BoolP("version", "V", false, "print version number")
}
