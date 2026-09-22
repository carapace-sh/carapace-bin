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

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"anchor-version": carapace.ActionValuesDescribed(
			"v1", "Generate Anchor v1 templates",
			"v2", "Generate Anchor v2 templates",
		),
		"package-manager": carapace.ActionValuesDescribed(
			"npm", "Use npm as the package manager",
			"yarn", "Use yarn as the package manager",
			"pnpm", "Use pnpm as the package manager",
			"bun", "Use bun as the package manager",
		),
		"template": carapace.ActionValuesDescribed(
			"single", "Program with a single `lib.rs` file (not recommended for production)",
			"multiple", "Program with multiple files for instructions, state... (recommended)",
		),
		"test-template": carapace.ActionValuesDescribed(
			"mocha", "Generate template for Mocha unit-test",
			"jest", "Generate template for Jest unit-test",
			"rust", "Generate template for Rust unit-test",
			"mollusk", "Generate template for Mollusk Rust unit-test",
			"litesvm", "Generate template for LiteSVM rust unit-test",
		),
	})
}
