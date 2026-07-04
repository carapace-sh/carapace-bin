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
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"create-keychain", "Create a new keychain",
			"delete-keychain", "Delete a keychain",
			"list-keychains", "List keychains",
			"show-keychain-info", "Show keychain information",
			"set-keychain-settings", "Set keychain settings",
			"lock-keychain", "Lock a keychain",
			"unlock-keychain", "Unlock a keychain",
			"import", "Import items into a keychain",
			"export", "Export items from a keychain",
			"find-certificate", "Find a certificate",
			"find-key", "Find a key",
			"find-identity", "Find an identity",
			"find-generic-password", "Find a generic password",
			"find-internet-password", "Find an internet password",
			"add-generic-password", "Add a generic password",
			"add-internet-password", "Add an internet password",
			"add-certificates", "Add certificates to a keychain",
			"delete-generic-password", "Delete a generic password",
			"delete-internet-password", "Delete an internet password",
			"delete-certificate", "Delete a certificate",
			"set-generic-password-partition-list", "Set partition list for a generic password",
			"dump-keychain", "Dump keychain contents",
			"authorizationdb", "Manipulate authorization database",
			"authorize", "Perform an authorization",
			"trust-settings", "Manage trust settings",
			"cms", "Create or verify CMS messages",
			"encode", "Encode/decode data",
			"authorizationdb", "Authorization database operations",
		),
	)
}
