package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Update repo with changes made in the underlying Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_importCmd).Standalone()

	git_importCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_importCmd)
}
