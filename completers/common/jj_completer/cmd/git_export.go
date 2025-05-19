package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Update the underlying Git repo with changes made in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_exportCmd).Standalone()

	git_exportCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_exportCmd)
}
