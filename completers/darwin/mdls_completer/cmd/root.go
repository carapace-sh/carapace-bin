package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdls",
	Short: "list Spotlight metadata attributes",
	Long:  "https://keith.github.io/xcode-manpages/mdls.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("attr", "n", "", "Fetch the specified attribute")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("nullMarker", "m", "", "Substitute string for null attributes in raw mode")
	rootCmd.Flags().StringP("plist", "p", "", "Output attributes in XML format to file")
	rootCmd.Flags().BoolP("raw", "r", false, "Print raw attribute data")
	rootCmd.Flags().BoolP("sdb", "s", false, "Get kMDItemSDBInfo attribute")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"plist": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
