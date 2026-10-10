package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var identifyCmd = &cobra.Command{
	Use:     "identify",
	Short:   "identify the working directory or specified revision",
	Aliases: []string{"id"},
	GroupID: groups[group_change_navigation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identifyCmd).Standalone()

	identifyCmd.Flags().BoolP("bookmarks", "B", false, "show bookmarks")
	identifyCmd.Flags().BoolP("branch", "b", false, "show branch")
	identifyCmd.Flags().BoolP("id", "i", false, "show global revision id")
	identifyCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	identifyCmd.Flags().BoolP("num", "n", false, "show local revision number")
	identifyCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	identifyCmd.Flags().StringP("rev", "r", "", "identify the specified revision")
	identifyCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	identifyCmd.Flags().BoolP("tags", "t", false, "show tags")
	identifyCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(identifyCmd)

	carapace.Gen(identifyCmd).FlagCompletion(carapace.ActionMap{
		"branch":    hg.ActionBranches(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       hg.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(identifyCmd).PositionalCompletion(
		carapace.Batch(hg.ActionPaths(), carapace.ActionFiles()).ToA(),
	)
}
