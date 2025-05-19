package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installBundlerGemsCmd = &cobra.Command{
	Use:     "install-bundler-gems",
	Short:   "Install Homebrew's Bundler gems",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installBundlerGemsCmd).Standalone()

	installBundlerGemsCmd.Flags().Bool("add-groups", false, "Installs the specified comma-separated list of gem groups, in addition to those already installed.")
	installBundlerGemsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	installBundlerGemsCmd.Flags().Bool("groups", false, "Installs the specified comma-separated list of gem groups (default: last used). Replaces any previously installed groups.")
	installBundlerGemsCmd.Flags().Bool("help", false, "Show this message.")
	installBundlerGemsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	installBundlerGemsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(installBundlerGemsCmd)
}
