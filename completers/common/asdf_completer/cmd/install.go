package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/asdf"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a specific version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("keep-download", false, "Whether or not to keep download directory after successful install")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalCompletion(
		asdf.ActionInstalledPlugins(),
		asdf.ActionInstallVersions(""),
	)
}
