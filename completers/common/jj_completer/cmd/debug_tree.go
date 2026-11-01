package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var debug_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "List the recursive entries of a tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_treeCmd).Standalone()

	debug_treeCmd.Flags().String("dir", "", "")
	debug_treeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_treeCmd.Flags().String("id", "", "")
	debug_treeCmd.Flags().StringP("revision", "r", "", "")
	debugCmd.AddCommand(debug_treeCmd)

	carapace.Gen(debug_treeCmd).FlagCompletion(carapace.ActionMap{
		"id":       jj.ActionRevsets(jj.RevOpts{}.Default()),
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})
}
