package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var env_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all Environment Variables for a Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_lsCmd).Standalone()

	env_lsCmd.Flags().String("format", "", "Output format")
	env_lsCmd.Flags().String("guidance", "", "Receive command suggestions")
	env_lsCmd.Flags().Bool("json", false, "Output as JSON")
	env_lsCmd.Flags().String("project", "", "Project name or ID")

	envCmd.AddCommand(env_lsCmd)

	carapace.Gen(env_lsCmd).PositionalCompletion(
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
