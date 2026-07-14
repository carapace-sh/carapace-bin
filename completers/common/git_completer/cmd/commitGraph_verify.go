package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commitGraph_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Read the commit-graph file and verify its contents against the object database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitGraph_verifyCmd).Standalone()

	commitGraph_verifyCmd.Flags().Bool("no-progress", false, "turn progress off explicitly")
	commitGraph_verifyCmd.Flags().String("object-dir", "", "use given directory for the location of packfiles and commit-graph file")
	commitGraph_verifyCmd.Flags().Bool("progress", false, "turn progress on explicitly")
	commitGraph_verifyCmd.Flags().Bool("shallow", false, "only check the tip commit-graph file in a chain of split commit-graphs")
	commitGraphCmd.AddCommand(commitGraph_verifyCmd)

	carapace.Gen(commitGraph_verifyCmd).FlagCompletion(carapace.ActionMap{
		"object-dir": carapace.ActionDirectories(),
	})
}
