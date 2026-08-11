package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff two branches using the common ancestor base revision Will calculate the set of changes between source branch latest revision and the base revision that is not in the set of changes between the target branch latest revision and the base revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_diffCmd).Standalone()

	branch_diffCmd.Flags().Bool("auto-resolve", false, "Attempt to auto resolve conflicts if true")
	branch_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_diffCmd.Flags().String("source", "", "Name of the source branch")
	branchCmd.AddCommand(branch_diffCmd)
}
