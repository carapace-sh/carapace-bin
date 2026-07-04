package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codesign",
	Short: "create and manipulate code signatures",
	Long:  "https://keith.github.io/xcode-manpages/codesign.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("display", "d", false, "Display the code signature")
	rootCmd.Flags().BoolP("entitlements", "e", false, "Print the entitlements")
	rootCmd.Flags().BoolP("force", "f", false, "Replace any existing signature")
	rootCmd.Flags().StringP("identifier", "i", "", "Identifier to use for code signing")
	rootCmd.Flags().StringP("requirements", "r", "", "Requirements for signing or verification")
	rootCmd.Flags().StringP("sign", "s", "", "Sign the code with the given identity")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode (set multiple times for more verbosity)")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
