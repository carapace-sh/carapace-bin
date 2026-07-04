package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "security",
	Short: "keychain and certificate management",
	Long:  "https://keith.github.io/xcode-manpages/security.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode")
	rootCmd.Flags().BoolP("leaks", "l", false, "Run leaks on exit")
	rootCmd.Flags().StringP("prompt", "p", "", "Set the interactive prompt")
	rootCmd.Flags().BoolP("quiet", "q", false, "Less verbose output")
	rootCmd.Flags().BoolP("verbose", "v", false, "More verbose output")
}
