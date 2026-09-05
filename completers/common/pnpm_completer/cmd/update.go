package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update packages to their newest version based on the specified range",
	Aliases: []string{"up", "upgrade"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("changeset", false, "Generate a changeset file declaring a patch bump for every workspace package whose production dependencies were changed by the update")
	updateCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	updateCmd.Flags().String("depth", "", "How deep to inspect dependencies. `0` means top-level dependencies only. Defaults to unlimited")
	updateCmd.Flags().BoolP("dev", "D", false, "Update packages only in \"devDependencies\"")
	updateCmd.Flags().BoolP("global", "g", false, "Update globally installed packages")
	updateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	updateCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	updateCmd.Flags().Bool("include-github-actions", false, "Also update GitHub Actions dependencies in workflow and action files")
	updateCmd.Flags().BoolP("interactive", "i", false, "Show outdated dependencies and select which ones to update")
	updateCmd.Flags().BoolP("latest", "L", false, "Ignore version ranges in package.json: bump the matched packages to their latest version and rewrite the manifest ranges")
	updateCmd.Flags().StringSlice("libc", nil, "libc families whose platform-specific optional dependencies should be installed (`glibc`, `musl`). Repeat or comma-separate for multiple values")
	updateCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	updateCmd.Flags().Bool("lockfile-only", false, "Dependencies are not downloaded; only `pnpm-lock.yaml` is updated")
	updateCmd.Flags().Bool("no-changeset", false, "Do not generate a changeset, even when `updateConfig.changeset` enables changeset generation by default")
	updateCmd.Flags().Bool("no-optional", false, "Don't update packages in \"optionalDependencies\"")
	updateCmd.Flags().Bool("no-save", false, "Do not write the updated ranges back to package.json. The lockfile is still updated (the `--no-save` flag)")
	updateCmd.Flags().Bool("optional", false, "Update packages only in \"optionalDependencies\"")
	updateCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	updateCmd.Flags().Bool("patches", false, "Refresh registry revisions without changing package versions")
	updateCmd.Flags().String("pnpr-server", "", "URL of a pnpr server to offload revision refresh resolution to")
	updateCmd.Flags().BoolP("prod", "P", false, "Update packages only in \"dependencies\" and \"optionalDependencies\"")
	updateCmd.Flags().Bool("production", false, "Update packages only in \"dependencies\" and \"optionalDependencies\"")
	updateCmd.Flags().BoolP("save-exact", "E", false, "Write the resolved version without a range operator when rewriting the manifest under `--latest`")
	updateCmd.Flags().Bool("workspace", false, "Tries to link all packages from the workspace, updating versions to match the workspace packages")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
