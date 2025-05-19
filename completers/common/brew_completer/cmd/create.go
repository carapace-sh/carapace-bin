package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Generate a formula or, with `--cask`, a cask for the downloadable file at <URL> and open it in the editor",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("HEAD", false, "Indicate that <URL> points to the package's repository rather than a file.")
	createCmd.Flags().Bool("autotools", false, "Create a basic template for an Autotools-style build.")
	createCmd.Flags().Bool("cask", false, "Create a basic template for a cask.")
	createCmd.Flags().Bool("cmake", false, "Create a basic template for a CMake-style build.")
	createCmd.Flags().Bool("crystal", false, "Create a basic template for a Crystal build.")
	createCmd.Flags().Bool("debug", false, "Display any debugging information.")
	createCmd.Flags().Bool("force", false, "Ignore errors for disallowed formula names and names that shadow aliases.")
	createCmd.Flags().Bool("go", false, "Create a basic template for a Go build.")
	createCmd.Flags().Bool("help", false, "Show this message.")
	createCmd.Flags().Bool("meson", false, "Create a basic template for a Meson-style build.")
	createCmd.Flags().Bool("no-fetch", false, "Homebrew will not download <URL> to the cache and will thus not add its SHA-256 to the formula for you, nor will it check the GitHub API for GitHub projects (to fill out its description and homepage).")
	createCmd.Flags().Bool("node", false, "Create a basic template for a Node build.")
	createCmd.Flags().Bool("perl", false, "Create a basic template for a Perl build.")
	createCmd.Flags().Bool("python", false, "Create a basic template for a Python build.")
	createCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	createCmd.Flags().Bool("ruby", false, "Create a basic template for a Ruby build.")
	createCmd.Flags().Bool("rust", false, "Create a basic template for a Rust build.")
	createCmd.Flags().Bool("set-license", false, "Explicitly set the <license> of the new formula.")
	createCmd.Flags().Bool("set-name", false, "Explicitly set the <name> of the new formula or cask.")
	createCmd.Flags().Bool("set-version", false, "Explicitly set the <version> of the new formula or cask.")
	createCmd.Flags().Bool("tap", false, "Generate the new formula within the given tap, specified as <user>`/`<repo>.")
	createCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(createCmd)
}
