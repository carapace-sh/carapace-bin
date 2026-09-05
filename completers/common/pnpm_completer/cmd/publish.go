package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a package to the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().String("access", "", "Publish the package as `public` or `restricted`")
	publishCmd.Flags().Bool("batch", false, "Send all workspace packages in a single request (requires `--recursive`)")
	publishCmd.Flags().Bool("dry-run", false, "Do everything `publish` would do except uploading to the registry")
	publishCmd.Flags().Bool("embed-readme", false, "Embed the README contents in the published manifest")
	publishCmd.Flags().Bool("force", false, "Publish even if the version is already in the registry")
	publishCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	publishCmd.Flags().Bool("ignore-scripts", false, "Don't run publish-related lifecycle scripts")
	publishCmd.Flags().Bool("json", false, "Print the per-package publish summary in JSON")
	publishCmd.Flags().Bool("no-embed-readme", false, "Do not embed README contents in the published manifest")
	publishCmd.Flags().Bool("no-git-checks", false, "Skip the git working-tree / branch / remote checks")
	publishCmd.Flags().Bool("no-skip-manifest-obfuscation", false, "Apply pnpm's normal published-manifest filtering")
	publishCmd.Flags().String("otp", "", "One-time password for two-factor-authenticated registries")
	publishCmd.Flags().Bool("provenance", false, "Generate a provenance attestation for the published package")
	publishCmd.Flags().String("publish-branch", "", "The branch publishing is allowed from. Defaults to `master` / `main`")
	publishCmd.Flags().Bool("report-summary", false, "Recursive only: write a `pnpm-publish-summary.json` report listing the packages that were published")
	publishCmd.Flags().Bool("skip-manifest-obfuscation", false, "Keep the original `packageManager` field and publish-lifecycle scripts in the published manifest instead of stripping them")
	publishCmd.Flags().String("tag", "", "Register the published package under this tag instead of `latest`")
	publishCmd.Flag("no-embed-readme").Hidden = true
	publishCmd.Flag("no-skip-manifest-obfuscation").Hidden = true

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"access": carapace.ActionValues("public", "restricted"),
	})

	rootCmd.AddCommand(publishCmd)
}
