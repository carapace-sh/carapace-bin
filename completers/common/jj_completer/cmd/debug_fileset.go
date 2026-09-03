package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var debug_filesetCmd = &cobra.Command{
	Use:   "fileset",
	Short: "Parse fileset expression",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_filesetCmd).Standalone()

	debug_filesetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_filesetCmd)

	carapace.Gen(debug_filesetCmd).PositionalCompletion(
		jj.ActionFilesets(),
	)
}