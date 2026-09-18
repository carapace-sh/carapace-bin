package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Apply the pending change intents (`pnpm version -r`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("allow-same-version", false, "Allow bumping to the same version")
	versionCmd.Flags().Bool("dry-run", false, "Print what the command would do without changing anything")
	versionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	versionCmd.Flags().Bool("json", false, "Show information in JSON format")
	versionCmd.Flags().String("message", "", "Commit message. \"%s\" is replaced with the new version. Default is \"%s\"")
	versionCmd.Flags().Bool("no-commit-hooks", false, "Skip running git commit hooks when committing the version bump")
	versionCmd.Flags().Bool("no-git-checks", false, "Don't check if the working tree is clean")
	versionCmd.Flags().Bool("no-git-tag-version", false, "Don't create a commit or tag for the version bump. Git commits and tags are always skipped in recursive mode")
	versionCmd.Flags().String("preid", "", "Sets the prerelease identifier (e.g. alpha, beta, rc)")
	versionCmd.Flags().Bool("sign-git-tag", false, "Sign the generated git tag with GPG")
	versionCmd.Flags().String("tag-version-prefix", "v", "Sets the tag prefix. Default is \"v\". Set to empty string to remove the prefix")
	rootCmd.AddCommand(versionCmd)
}
