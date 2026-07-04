package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "manage-bde",
	Short: "manage BitLocker Drive Encryption",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/manage-bde",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"status", "display BitLocker status",
			"on", "enable BitLocker on a volume",
			"off", "disable BitLocker on a volume",
			"pause", "pause BitLocker on a volume",
			"resume", "resume BitLocker on a volume",
			"lock", "prevent access to BitLocker-encrypted data",
			"unlock", "allow access to BitLocker-encrypted data",
			"autounlock", "manage automatic unlocking",
			"protectors", "manage protection methods",
			"key", "manage BitLocker keys",
			"upgrade", "upgrade BitLocker",
			"wipe", "wipe free space",
		),
	)
}
