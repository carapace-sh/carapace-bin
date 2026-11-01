package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update dependencies listed in kcl.mod.lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_updateCmd).Standalone()

	mod_updateCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the KCL module download")
	mod_updateCmd.Flags().Bool("no_sum_check", false, "do not check the checksum of the package and update kcl.mod.lock")
	modCmd.AddCommand(mod_updateCmd)
}
