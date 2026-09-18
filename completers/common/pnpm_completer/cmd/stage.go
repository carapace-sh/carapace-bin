package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage packages for publishing, deferring proof-of-presence (2FA) to a later point in time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stageCmd).Standalone()

	stageCmd.Flags().String("access", "", "Publish the package as `public` or `restricted`")
	stageCmd.Flags().Bool("batch", false, "Send all workspace packages in a single request (requires `--recursive`)")
	stageCmd.Flags().Bool("dry-run", false, "Do everything `publish` would do except uploading to the registry")
	stageCmd.Flags().Bool("embed-readme", false, "Embed the README contents in the published manifest")
	stageCmd.Flags().Bool("force", false, "Publish even if the version is already in the registry")
	stageCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stageCmd.Flags().Bool("ignore-scripts", false, "Don't run publish-related lifecycle scripts")
	stageCmd.Flags().Bool("json", false, "Print the per-package publish summary in JSON")
	stageCmd.Flags().Bool("no-embed-readme", false, "Do not embed README contents in the published manifest")
	stageCmd.Flags().Bool("no-git-checks", false, "Skip the git working-tree / branch / remote checks")
	stageCmd.Flags().Bool("no-skip-manifest-obfuscation", false, "Apply pnpm's normal published-manifest filtering")
	stageCmd.Flags().String("otp", "", "One-time password for two-factor-authenticated registries")
	stageCmd.Flags().Bool("provenance", false, "Generate a provenance attestation for the published package")
	stageCmd.Flags().String("publish-branch", "", "The branch publishing is allowed from. Defaults to `master` / `main`")
	stageCmd.Flags().String("registry", "", "The base URL of the npm registry")
	stageCmd.Flags().Bool("report-summary", false, "Recursive only: write a `pnpm-publish-summary.json` report listing the packages that were published")
	stageCmd.Flags().Bool("skip-manifest-obfuscation", false, "Keep the original `packageManager` field and publish-lifecycle scripts in the published manifest instead of stripping them")
	stageCmd.Flags().String("tag", "", "Register the published package under this tag instead of `latest`")
	stageCmd.Flag("no-embed-readme").Hidden = true
	stageCmd.Flag("no-skip-manifest-obfuscation").Hidden = true

	carapace.Gen(stageCmd).FlagCompletion(carapace.ActionMap{
		"access": carapace.ActionValues("public", "restricted"),
	})

	rootCmd.AddCommand(stageCmd)
}
