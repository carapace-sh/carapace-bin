package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_mirrorCmd = &cobra.Command{
	Use:   "mirror [ID | URL | PATH] [flags]",
	Short: "Mirror a project or repository to the specified location, using pull or push methods.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_mirrorCmd).Standalone()

	repo_mirrorCmd.Flags().Bool("allow-divergence", false, "Determines if divergent refs are skipped.")
	repo_mirrorCmd.Flags().String("direction", "", "Mirror direction. Options: pull, push.")
	repo_mirrorCmd.Flags().Bool("enabled", false, "Determines if the mirror is enabled.")
	repo_mirrorCmd.Flags().Bool("protected-branches-only", false, "Determines if only protected branches are mirrored.")
	repo_mirrorCmd.Flags().String("url", "", "The target URL to which the repository is mirrored.")
	repo_mirrorCmd.MarkFlagRequired("direction")
	repo_mirrorCmd.MarkFlagRequired("url")
	repoCmd.AddCommand(repo_mirrorCmd)

	carapace.Gen(repo_mirrorCmd).FlagCompletion(carapace.ActionMap{
		"direction": carapace.ActionValues("pull", "push"),
	})

	carapace.Gen(repo_mirrorCmd).PositionalCompletion(
		action.ActionRepo(repo_mirrorCmd),
	)
}
