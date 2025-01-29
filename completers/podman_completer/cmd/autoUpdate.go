package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoUpdateCmd = &cobra.Command{
	Use:   "auto-update [options]",
	Short: "Auto update containers according to their auto-update policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoUpdateCmd).Standalone()

	autoUpdateCmd.Flags().String("authfile", "", "Path to the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	autoUpdateCmd.Flags().Bool("dry-run", false, "Check for pending updates")
	autoUpdateCmd.Flags().String("format", "", "Change the output format to JSON or a Go template")
	autoUpdateCmd.Flags().Bool("rollback", false, "Rollback to previous image if update fails")
	autoUpdateCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	rootCmd.AddCommand(autoUpdateCmd)
}
