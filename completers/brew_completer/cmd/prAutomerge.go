package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prAutomergeCmd = &cobra.Command{
	Use:     "pr-automerge",
	Short:   "Find pull requests that can be automatically merged using `brew pr-publish`",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prAutomergeCmd).Standalone()

	prAutomergeCmd.Flags().Bool("autosquash", false, "Instruct `brew pr-publish` to automatically reformat and reword commits in the pull request to the preferred format.")
	prAutomergeCmd.Flags().Bool("debug", false, "Display any debugging information.")
	prAutomergeCmd.Flags().Bool("help", false, "Show this message.")
	prAutomergeCmd.Flags().Bool("ignore-failures", false, "Include pull requests that have failing status checks.")
	prAutomergeCmd.Flags().Bool("publish", false, "Run `brew pr-publish` on matching pull requests.")
	prAutomergeCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	prAutomergeCmd.Flags().Bool("tap", false, "Target tap repository (default: `homebrew/core`).")
	prAutomergeCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	prAutomergeCmd.Flags().Bool("with-label", false, "Pull requests must have this label.")
	prAutomergeCmd.Flags().Bool("without-approval", false, "Pull requests do not require approval to be merged.")
	prAutomergeCmd.Flags().Bool("without-labels", false, "Pull requests must not have these labels (default: `do not merge`, `new formula`, `automerge-skip`, `pre-release`, `CI-published-bottle-commits`).")
	prAutomergeCmd.Flags().Bool("workflow", false, "Workflow file to use with `brew pr-publish`.")
	rootCmd.AddCommand(prAutomergeCmd)
}
