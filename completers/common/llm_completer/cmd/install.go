package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install packages from PyPI into the same environment as LLM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringP("editable", "e", "", "Install a project in editable mode from this path")
	installCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages even if they are already up-to-date")
	installCmd.Flags().Bool("no-cache-dir", false, "Disable the cache")
	installCmd.Flags().Bool("pre", false, "Include pre-release and development versions")
	installCmd.Flags().BoolP("upgrade", "U", false, "Upgrade packages to latest version")
	rootCmd.AddCommand(installCmd)
}
