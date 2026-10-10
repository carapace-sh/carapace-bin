package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugbuilddagCmd = &cobra.Command{
	Use:    "debugbuilddag",
	Short:  "builds a repo with a given DAG from scratch in the current empty repo",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugbuilddagCmd).Standalone()

	debugbuilddagCmd.Flags().Bool("from-existing", false, "continue from a non-empty repository")
	debugbuilddagCmd.Flags().BoolP("mergeable-file", "m", false, "add single file mergeable changes")
	debugbuilddagCmd.Flags().BoolP("new-file", "n", false, "add new file at each rev")
	debugbuilddagCmd.Flags().BoolP("overwritten-file", "o", false, "add single file all revs overwrite")
	rootCmd.AddCommand(debugbuilddagCmd)
}
