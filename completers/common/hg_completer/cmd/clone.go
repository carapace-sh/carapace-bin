package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Short:   "make a copy of an existing repository",
	GroupID: groups[group_repository_creation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().StringArrayP("branch", "b", nil, "do not clone everything, but include this branch's changesets and their ancestors")
	cloneCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	cloneCmd.Flags().BoolP("noupdate", "U", false, "the clone will include an empty working directory (only a repository)")
	cloneCmd.Flags().Bool("pull", false, "use pull protocol to copy metadata")
	cloneCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	cloneCmd.Flags().StringArrayP("rev", "r", nil, "do not clone everything, but include this changeset and its ancestors")
	cloneCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	cloneCmd.Flags().Bool("stream", false, "clone with minimal data processing")
	cloneCmd.Flags().Bool("uncompressed", false, "an alias to --stream (DEPRECATED)")
	cloneCmd.Flags().StringP("updaterev", "u", "", "revision, tag, or branch to check out")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"branch":    action.ActionBranches(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       action.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		carapace.Batch(action.ActionPaths(), carapace.ActionFiles()).ToA(),
		carapace.ActionDirectories(),
	)
}
