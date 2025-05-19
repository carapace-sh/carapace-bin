package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tapInfoCmd = &cobra.Command{
	Use:     "tap-info",
	Short:   "Show detailed information about one or more <tap>s",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tapInfoCmd).Standalone()

	tapInfoCmd.Flags().Bool("debug", false, "Display any debugging information.")
	tapInfoCmd.Flags().Bool("help", false, "Show this message.")
	tapInfoCmd.Flags().Bool("installed", false, "Show information on each installed tap.")
	tapInfoCmd.Flags().Bool("json", false, "Print a JSON representation of <tap>. Currently the default and only accepted value for <version> is `v1`. See the docs for examples of using the JSON output: <https://docs.brew.sh/Querying-Brew>")
	tapInfoCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	tapInfoCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(tapInfoCmd)
}
