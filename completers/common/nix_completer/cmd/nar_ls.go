package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var nar_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "show information about a path inside a NAR file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nar_lsCmd).Standalone()

	nar_lsCmd.Flags().BoolP("directory", "d", false, "Show directories rather than their contents")
	nar_lsCmd.Flags().Bool("json", false, "Produce output in JSON format")
	nar_lsCmd.Flags().BoolP("long", "l", false, "Show detailed file information")
	nar_lsCmd.Flags().BoolP("recursive", "R", false, "List subdirectories recursively")
	narCmd.AddCommand(nar_lsCmd)

	addLoggingFlags(nar_lsCmd)

	carapace.Gen(nar_lsCmd).PositionalCompletion(
		carapace.ActionFiles(".nar"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return nix.ActionNarFileContents(c.Args[0])
		}),
	)
}
