package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("anchor-version", "v1", "Anchor template version to generate")
	initCmd.Flags().Bool("force", false, "Initialize even if there are files")
	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	initCmd.Flags().Bool("install-agent-skills", false, "Install Solana agent skills")
	initCmd.Flags().BoolP("javascript", "j", false, "Use JavaScript instead of TypeScript")
	initCmd.Flags().Bool("no-git", false, "Don't initialize git")
	initCmd.Flags().Bool("no-install", false, "Don't install JavaScript dependencies")
	initCmd.Flags().String("package-manager", "", "Package Manager to use. If omitted, detection cascades `pnpm` -> `yarn` -> `npm` and picks the first one on PATH. When set explicitly, the chosen binary must be installed")
	initCmd.Flags().StringP("template", "t", "multiple", "Rust program template to use")
	initCmd.Flags().String("test-template", "litesvm", "Test template to use")
	rootCmd.AddCommand(initCmd)
}
