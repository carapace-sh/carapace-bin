package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var patchIdCmd = &cobra.Command{
	Use:     "patch-id",
	Short:   "Print stable patch-IDs for one or more commits",
	GroupID: "plumbing",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchIdCmd).Standalone()

	patchIdCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	patchIdCmd.Flags().Bool("unstable", false, "Use legacy unstable patch-ID algorithm (Git 1.9 and older)")
	rootCmd.AddCommand(patchIdCmd)

	carapace.Gen(patchIdCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
