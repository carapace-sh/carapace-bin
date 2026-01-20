package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var rewordCmd = &cobra.Command{
	Use:   "reword",
	Short: "Edit the commit message of the specified commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewordCmd).Standalone()

	rewordCmd.Flags().BoolP("format", "f", false, "Format the existing commit message to 72-char line wrapping without opening an editor")
	rewordCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rewordCmd.Flags().StringP("message", "m", "", "The new commit message or branch name. If not provided, opens an editor")
	rootCmd.AddCommand(rewordCmd)

	carapace.Gen(rewordCmd).PositionalCompletion(
		but.ActionTargets(),
	)
}
