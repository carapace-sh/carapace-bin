package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpTrustSettingsCmd = &cobra.Command{
	Use:   "dump-trust-settings",
	Short: "Display contents of trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpTrustSettingsCmd).Standalone()
	dumpTrustSettingsCmd.Flags().BoolP("admin", "d", false, "Display trusted admin certs (default: user)")
	dumpTrustSettingsCmd.Flags().BoolP("system", "s", false, "Display trusted system certs (default: user)")
	rootCmd.AddCommand(dumpTrustSettingsCmd)
}
