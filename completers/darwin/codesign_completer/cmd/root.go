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

	rootCmd.Flags().StringP("bundle-identifier", "bundle-identifier", "", "Identifier to use for code signing")
	rootCmd.Flags().StringP("certificate", "certificate", "", "Use the specified certificate")
	rootCmd.Flags().BoolP("deep", "deep", false, "Sign all nested code")
	rootCmd.Flags().BoolP("display", "display", false, "Display the code signature")
	rootCmd.Flags().BoolP("entitlements", "entitlements", false, "Print the entitlements")
	rootCmd.Flags().BoolP("force", "force", false, "Replace any existing signature")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("requirements", "requirements", false, "Print the requirements")
	rootCmd.Flags().StringP("resource-rules", "resource-rules", "", "Use the specified resource rules plist")
	rootCmd.Flags().BoolP("sign", "sign", false, "Sign the code with the given identity")
	rootCmd.Flags().BoolP("strict", "strict", false, "Check the signature strictly")
	rootCmd.Flags().BoolP("verbose", "verbose", false, "Verbose mode")
	rootCmd.Flags().BoolP("verify", "verify", false, "Verify the code signature")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
