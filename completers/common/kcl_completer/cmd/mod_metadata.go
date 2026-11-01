package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Package metadata of a kcl package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_metadataCmd).Standalone()

	mod_metadataCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the KCL module download")
	mod_metadataCmd.Flags().Bool("update", false, "check the local package and update and download the local package. (default: false)")
	mod_metadataCmd.Flags().Bool("vendor", false, "run in vendor mode (default: false)")
	modCmd.AddCommand(mod_metadataCmd)
}
