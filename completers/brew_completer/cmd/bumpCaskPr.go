package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpCaskPrCmd = &cobra.Command{
	Use:     "bump-cask-pr",
	Short:   "Create a pull request to update <cask> with a new version",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpCaskPrCmd).Standalone()

	bumpCaskPrCmd.Flags().Bool("commit", false, "When passed with `--write-only`, generate a new commit after writing changes to the cask file.")
	bumpCaskPrCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bumpCaskPrCmd.Flags().Bool("dry-run", false, "Print what would be done rather than doing it.")
	bumpCaskPrCmd.Flags().Bool("force", false, "Ignore duplicate open PRs.")
	bumpCaskPrCmd.Flags().Bool("fork-org", false, "Use the specified GitHub organization for forking.")
	bumpCaskPrCmd.Flags().Bool("help", false, "Show this message.")
	bumpCaskPrCmd.Flags().Bool("message", false, "Prepend <message> to the default pull request message.")
	bumpCaskPrCmd.Flags().Bool("no-audit", false, "Don't run `brew audit` before opening the PR.")
	bumpCaskPrCmd.Flags().Bool("no-browse", false, "Print the pull request URL instead of opening in a browser.")
	bumpCaskPrCmd.Flags().Bool("no-fork", false, "Don't try to fork the repository.")
	bumpCaskPrCmd.Flags().Bool("no-style", false, "Don't run `brew style --fix` before opening the PR.")
	bumpCaskPrCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bumpCaskPrCmd.Flags().Bool("sha256", false, "Specify the <SHA-256> checksum of the new download.")
	bumpCaskPrCmd.Flags().Bool("url", false, "Specify the <URL> for the new download.")
	bumpCaskPrCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bumpCaskPrCmd.Flags().Bool("version", false, "Specify the new <version> for the cask.")
	bumpCaskPrCmd.Flags().Bool("version-arm", false, "Specify the new cask <version> for the ARM architecture.")
	bumpCaskPrCmd.Flags().Bool("version-intel", false, "Specify the new cask <version> for the Intel architecture.")
	bumpCaskPrCmd.Flags().Bool("write-only", false, "Make the expected file modifications without taking any Git actions.")
	rootCmd.AddCommand(bumpCaskPrCmd)
}
