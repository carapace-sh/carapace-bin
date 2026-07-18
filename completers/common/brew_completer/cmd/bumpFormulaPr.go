package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpFormulaPrCmd = &cobra.Command{
	Use:     "bump-formula-pr",
	Short:   "Create a pull request to update <formula> with a new URL or a new tag",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpFormulaPrCmd).Standalone()

	bumpFormulaPrCmd.Flags().Bool("commit", false, "When passed with `--write-only`, generate a new commit after writing changes to the formula file.")
	bumpFormulaPrCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bumpFormulaPrCmd.Flags().Bool("dry-run", false, "Print what would be done rather than doing it.")
	bumpFormulaPrCmd.Flags().Bool("force", false, "Ignore duplicate open PRs. Remove all mirrors if `--mirror` was not specified.")
	bumpFormulaPrCmd.Flags().String("fork-org", "", "Use the specified GitHub organization for forking.")
	bumpFormulaPrCmd.Flags().Bool("help", false, "Show this message.")
	bumpFormulaPrCmd.Flags().String("message", "", "Prepend <message> to the default pull request message.")
	bumpFormulaPrCmd.Flags().String("mirror", "", "Use the specified <URL> as a mirror URL. If <URL> is a comma-separated list of URLs, multiple mirrors will be added.")
	bumpFormulaPrCmd.Flags().Bool("no-audit", false, "Don't run `brew audit` before opening the PR.")
	bumpFormulaPrCmd.Flags().Bool("no-browse", false, "Print the pull request URL instead of opening in a browser.")
	bumpFormulaPrCmd.Flags().Bool("no-fork", false, "Don't try to fork the repository.")
	bumpFormulaPrCmd.Flags().Bool("online", false, "Run `brew audit --online` before opening the PR.")
	bumpFormulaPrCmd.Flags().String("python-exclude-packages", "", "Exclude these Python packages when finding resources.")
	bumpFormulaPrCmd.Flags().String("python-extra-packages", "", "Include these additional Python packages when finding resources.")
	bumpFormulaPrCmd.Flags().String("python-package-name", "", "Use the specified <package-name> when finding Python resources for <formula>. If no package name is specified, it will be inferred from the formula's stable URL.")
	bumpFormulaPrCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bumpFormulaPrCmd.Flags().String("revision", "", "Specify the new commit <revision> corresponding to the specified git <tag> or specified <version>.")
	bumpFormulaPrCmd.Flags().String("sha256", "", "Specify the <SHA-256> checksum of the new download.")
	bumpFormulaPrCmd.Flags().Bool("strict", false, "Run `brew audit --strict` before opening the PR.")
	bumpFormulaPrCmd.Flags().String("tag", "", "Specify the new git commit <tag> for the formula.")
	bumpFormulaPrCmd.Flags().String("url", "", "Specify the <URL> for the new download. If a <URL> is specified, the <SHA-256> checksum of the new download should also be specified.")
	bumpFormulaPrCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bumpFormulaPrCmd.Flags().String("version", "", "Use the specified <version> to override the value parsed from the URL or tag. Note that `--version=0` can be used to delete an existing version override from a formula if it has become redundant.")
	bumpFormulaPrCmd.Flags().Bool("write-only", false, "Make the expected file modifications without taking any Git actions.")
	rootCmd.AddCommand(bumpFormulaPrCmd)
}
