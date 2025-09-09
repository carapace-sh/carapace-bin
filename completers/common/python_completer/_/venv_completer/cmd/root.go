package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "venv",
	Short: "Creates virtual Python environments in one or more target directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("clear", false, "Delete the contents of the environment directory if it already exists, before environment creation.")
	rootCmd.Flags().Bool("copies", false, "Try to use copies rather than symlinks, even when symlinks are the default for the platform.")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().String("prompt", "", "Provides an alternative prompt prefix for this environment.")
	rootCmd.Flags().Bool("symlinks", false, "Try to use symlinks rather than copies, when symlinks are not the default for the platform.")
	rootCmd.Flags().Bool("system-site-packages", false, "Give the virtual environment access to the system site-packages dir.")
	rootCmd.Flags().Bool("upgrade", false, "Upgrade the environment directory to use this version of Python, assuming Python has been upgraded in-place.")
	rootCmd.Flags().Bool("upgrade-deps", false, "Upgrade core dependencies: pip setuptools to the latest version in PyPI")
	rootCmd.Flags().Bool("without-pip", false, "Skips installing or upgrading pip in the virtual environment (pip is bootstrapped by default)")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
