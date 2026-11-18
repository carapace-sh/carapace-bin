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
	installBundlerGemsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	installBundlerGemsCmd.Flags().Bool("groups", false, "Installs the specified comma-separated list of gem groups (default: last used). Replaces any previously installed groups.")
	installBundlerGemsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	installBundlerGemsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	installBundlerGemsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(installBundlerGemsCmd)
}
