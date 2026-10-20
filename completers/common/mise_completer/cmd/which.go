package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Show the path a tool's executable resolves to",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()

	whichCmd.Flags().Bool("complete", false, "")
	whichCmd.Flags().BoolP("help", "h", false, "Print help")
	whichCmd.Flags().Bool("plugin", false, "Show the name of the tool that provides the executable")
	whichCmd.Flags().StringP("tool", "t", "", "Look in a specific tool version, such as `mise which npm --tool=node@20`")
	whichCmd.Flags().Bool("version", false, "Show the version instead of the path")
	whichCmd.Flag("complete").Hidden = true
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).PositionalCompletion(
		mise.ActionBinNames(),
	)
}
