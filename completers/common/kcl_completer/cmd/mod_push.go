package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push kcl package to OCI registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_pushCmd).Standalone()

	mod_pushCmd.Flags().Bool("force", false, "force push the package to the registry")
	mod_pushCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the KCL module download")
	mod_pushCmd.Flags().String("tar_path", "", "packaged target path that will be pushed")
	mod_pushCmd.Flags().Bool("vendor", false, "run in vendor mode (default: false)")
	modCmd.AddCommand(mod_pushCmd)

	carapace.Gen(mod_pushCmd).FlagCompletion(carapace.ActionMap{
		"tar_path": carapace.ActionFiles(".tar"),
	})
}
