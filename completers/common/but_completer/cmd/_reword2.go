package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _reword2Cmd = &cobra.Command{
	Use:    "_reword2",
	Short:  "Edit a commit message or rename a branch",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_reword2Cmd).Standalone()

	_reword2Cmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	_reword2Cmd.Flags().BoolP("fix-formatting", "f", false, "Format the existing commit message to 72-character line wrapping without opening an editor")
	_reword2Cmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	_reword2Cmd.Flags().StringP("message", "m", "", "The new commit message or branch name. If omitted, an editor opens")
	rootCmd.AddCommand(_reword2Cmd)
}