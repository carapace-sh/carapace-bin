package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var httpFetchCmd = &cobra.Command{
	Use:     "http-fetch",
	Short:   "Download from a remote Git repository via HTTP",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(httpFetchCmd).Standalone()

	httpFetchCmd.Flags().String("index-pack-args", "", "command to run on the contents of the downloaded pack")
	httpFetchCmd.Flags().String("packfile", "", "fetch the packfile directly")
	httpFetchCmd.Flags().Bool("recover", false, "erify that everything reachable from target is fetched")
	httpFetchCmd.Flags().Bool("stdin", false, "expect commit-id lines on stdin")
	httpFetchCmd.Flags().BoolS("v", "v", false, "report what is downloaded")
	httpFetchCmd.Flags().StringS("w", "w", "", "write the commit-id into the specified filename")
	rootCmd.AddCommand(httpFetchCmd)

	httpFetchCmd.MarkFlagsMutuallyExclusive("packfile", "stdin")

	carapace.Gen(httpFetchCmd).FlagCompletion(carapace.ActionMap{
		"w": carapace.ActionFiles().Chdir("refs").ChdirF(traverse.GitDir),
	})
}
