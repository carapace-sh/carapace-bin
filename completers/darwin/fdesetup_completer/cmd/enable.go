package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable FileVault and optionally add user(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()

	enableCmd.Flags().String("certificate", "", "Path to a DER encoded certificate file")
	enableCmd.Flags().String("defer", "", "Defer enabling FileVault")
	enableCmd.Flags().String("delayminutes", "", "Number of minutes before restart")
	enableCmd.Flags().String("device", "", "FV context information location")
	enableCmd.Flags().Bool("dontaskatlogout", false, "Don't prompt at user logout")
	enableCmd.Flags().String("forceatlogin", "", "Force user to enable at login after n times")
	enableCmd.Flags().Bool("forcerestart", false, "Force restart after enabling")
	enableCmd.Flags().Bool("inputplist", false, "Read configuration from stdin")
	enableCmd.Flags().Bool("institutional", false, "Specify an institutional recovery key")
	enableCmd.Flags().String("key", "", "Path to authentication keychain private key file")
	enableCmd.Flags().Bool("keychain", false, "Use FileVaultMaster.keychain")
	enableCmd.Flags().Bool("norecoverykey", false, "Do not return a recovery key")
	enableCmd.Flags().Bool("outputplist", false, "Output key and computer info to stdout")
	enableCmd.Flags().Bool("personal", false, "Specify a personal recovery key")
	enableCmd.Flags().Bool("prompt", false, "Always prompt for authentication")
	enableCmd.Flags().Bool("quiet", false, "No status during enable")
	enableCmd.Flags().String("user", "", "Short user name")
	enableCmd.Flags().String("usertoadd", "", "Additional user name when enabling")
	enableCmd.Flags().String("uuid", "", "User UUID")
	enableCmd.Flags().Bool("verbose", false, "Enable verbose mode")

	carapace.Gen(enableCmd).FlagCompletion(carapace.ActionMap{
		"certificate": carapace.ActionFiles(),
		"defer":       carapace.ActionFiles(),
		"key":         carapace.ActionFiles(),
		"user":        os.ActionUsers(),
		"usertoadd":   os.ActionUsers(),
	})
}
