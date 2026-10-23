package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull kcl package from OCI registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_pullCmd).Standalone()

	mod_pullCmd.Flags().String("branch", "", "git repository branch")
	mod_pullCmd.Flags().String("commit", "", "git repository commit")
	mod_pullCmd.Flags().String("git", "", "git repository url")
	mod_pullCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the KCL module download")
	mod_pullCmd.Flags().String("oci", "", "oci repository url")
	mod_pullCmd.Flags().String("tag", "", "git or oci repository tag")
	modCmd.AddCommand(mod_pullCmd)
}
