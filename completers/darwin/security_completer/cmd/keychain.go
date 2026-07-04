package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var createKeychainCmd = &cobra.Command{
	Use:   "create-keychain",
	Short: "Create keychains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var defaultKeychainCmd = &cobra.Command{
	Use:   "default-keychain",
	Short: "Display or set the default keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deleteKeychainCmd = &cobra.Command{
	Use:   "delete-keychain",
	Short: "Delete keychains and remove them from the search list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dumpKeychainCmd = &cobra.Command{
	Use:   "dump-keychain",
	Short: "Dump the contents of one or more keychains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listKeychainsCmd = &cobra.Command{
	Use:   "list-keychains",
	Short: "Display or manipulate the keychain search list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var lockKeychainCmd = &cobra.Command{
	Use:   "lock-keychain",
	Short: "Lock a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var loginKeychainCmd = &cobra.Command{
	Use:   "login-keychain",
	Short: "Display or set the login keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var setKeychainPasswordCmd = &cobra.Command{
	Use:   "set-keychain-password",
	Short: "Set password for a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var setKeychainSettingsCmd = &cobra.Command{
	Use:   "set-keychain-settings",
	Short: "Set settings for a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var showKeychainInfoCmd = &cobra.Command{
	Use:   "show-keychain-info",
	Short: "Show the settings for a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var unlockKeychainCmd = &cobra.Command{
	Use:   "unlock-keychain",
	Short: "Unlock a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createKeychainCmd).Standalone()
	createKeychainCmd.Flags().StringP("password", "p", "", "Use password for the keychains being created")
	createKeychainCmd.Flags().BoolP("prompt-password", "P", false, "Prompt for password using SecurityAgent")
	rootCmd.AddCommand(createKeychainCmd)
	carapace.Gen(createKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(defaultKeychainCmd).Standalone()
	defaultKeychainCmd.Flags().StringP("domain", "d", "", "Use the specified preference domain")
	defaultKeychainCmd.Flags().StringP("set", "s", "", "Set the default keychain")
	rootCmd.AddCommand(defaultKeychainCmd)

	carapace.Gen(deleteKeychainCmd).Standalone()
	rootCmd.AddCommand(deleteKeychainCmd)
	carapace.Gen(deleteKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(dumpKeychainCmd).Standalone()
	dumpKeychainCmd.Flags().BoolP("acl", "a", false, "Dump access control list of items")
	dumpKeychainCmd.Flags().BoolP("data", "d", false, "Dump decrypted data of items")
	dumpKeychainCmd.Flags().BoolP("interactive", "i", false, "Interactive access control list editing mode")
	dumpKeychainCmd.Flags().BoolP("raw", "r", false, "Dump raw encrypted data of items")
	rootCmd.AddCommand(dumpKeychainCmd)
	carapace.Gen(dumpKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(listKeychainsCmd).Standalone()
	listKeychainsCmd.Flags().StringP("domain", "d", "", "Use the specified preference domain")
	listKeychainsCmd.Flags().StringP("set", "s", "", "Set the search list to the specified keychains")
	rootCmd.AddCommand(listKeychainsCmd)

	carapace.Gen(lockKeychainCmd).Standalone()
	lockKeychainCmd.Flags().BoolP("all", "a", false, "Lock all keychains")
	rootCmd.AddCommand(lockKeychainCmd)
	carapace.Gen(lockKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(loginKeychainCmd).Standalone()
	loginKeychainCmd.Flags().StringP("domain", "d", "", "Use the specified preference domain")
	loginKeychainCmd.Flags().StringP("set", "s", "", "Set the login keychain")
	rootCmd.AddCommand(loginKeychainCmd)

	carapace.Gen(setKeychainPasswordCmd).Standalone()
	setKeychainPasswordCmd.Flags().StringP("old-password", "o", "", "Old keychain password")
	setKeychainPasswordCmd.Flags().StringP("password", "p", "", "New keychain password")
	rootCmd.AddCommand(setKeychainPasswordCmd)
	carapace.Gen(setKeychainPasswordCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(setKeychainSettingsCmd).Standalone()
	setKeychainSettingsCmd.Flags().BoolP("lock-after-timeout", "u", false, "Lock keychain after timeout interval")
	setKeychainSettingsCmd.Flags().BoolP("lock-on-sleep", "l", false, "Lock keychain when the system sleeps")
	setKeychainSettingsCmd.Flags().StringP("timeout", "t", "", "Timeout interval in seconds (omit for no timeout)")
	rootCmd.AddCommand(setKeychainSettingsCmd)
	carapace.Gen(setKeychainSettingsCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(showKeychainInfoCmd).Standalone()
	rootCmd.AddCommand(showKeychainInfoCmd)
	carapace.Gen(showKeychainInfoCmd).PositionalAnyCompletion(security.ActionKeychains())

	carapace.Gen(unlockKeychainCmd).Standalone()
	unlockKeychainCmd.Flags().StringP("password", "p", "", "Use password to unlock the keychain")
	rootCmd.AddCommand(unlockKeychainCmd)
	carapace.Gen(unlockKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())
}
