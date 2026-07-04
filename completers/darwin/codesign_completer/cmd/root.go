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

	rootCmd.Flags().Bool("all-architectures", false, "Sign or verify all architectures")
	rootCmd.Flags().StringP("arch", "a", "", "Architecture to sign or verify")
	rootCmd.Flags().String("bundle-version", "", "Bundle version string")
	rootCmd.Flags().Bool("check-notarization", false, "Check notarization status")
	rootCmd.Flags().Bool("deep", false, "Recursively sign nested code (deprecated for signing)")
	rootCmd.Flags().StringP("detached", "D", "", "Use detached signature file")
	rootCmd.Flags().Bool("detached-database", false, "Use detached signature with database")
	rootCmd.Flags().BoolP("display", "d", false, "Display the code signature")
	rootCmd.Flags().String("entitlements", "", "Path to entitlements plist")
	rootCmd.Flags().BoolP("force", "f", false, "Replace any existing signature")
	rootCmd.Flags().Bool("force-library-entitlements", false, "Force library entitlements")
	rootCmd.Flags().Bool("generate-entitlement-der", false, "Generate entitlement DER")
	rootCmd.Flags().BoolP("hosting", "h", false, "Display hosting paths")
	rootCmd.Flags().StringP("identifier", "i", "", "Identifier to use for code signing")
	rootCmd.Flags().StringP("options", "o", "", "Signing options (flags comma-separated)")
	rootCmd.Flags().StringP("pagesize", "P", "", "Page size for detached signing")
	rootCmd.Flags().Bool("remove-signature", false, "Remove the existing signature")
	rootCmd.Flags().StringP("requirements", "r", "", "Requirements for signing or verification")
	rootCmd.Flags().StringP("sign", "s", "", "Sign the code with the given identity")
	rootCmd.Flags().StringP("test-requirement", "R", "", "Test requirement string or file")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode (set multiple times for more verbosity). Also used for verification mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"arch":         carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
		"detached":     carapace.ActionFiles(),
		"entitlements": carapace.ActionFiles(),
		"sign":         carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
