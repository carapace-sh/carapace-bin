package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_addCmd).Standalone()

	mod_addCmd.Flags().String("branch", "", "git repository branch")
	mod_addCmd.Flags().String("commit", "", "git repository commit")
	mod_addCmd.Flags().String("git", "", "git repository url")
	mod_addCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the KCL module download")
	mod_addCmd.Flags().Bool("no_sum_check", false, "do not check the checksum of the package and update kcl.mod.lock")
	mod_addCmd.Flags().String("oci", "", "oci repository url")
	mod_addCmd.Flags().String("path", "", "filesystem path to local dependency to add")
	mod_addCmd.Flags().String("rename", "", "rename the dependency")
	mod_addCmd.Flags().String("tag", "", "git or oci repository tag")
	modCmd.AddCommand(mod_addCmd)

	carapace.Gen(mod_addCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
}
