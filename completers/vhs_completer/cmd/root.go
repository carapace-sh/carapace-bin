package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vhs <file>",
	Short: "Run a given tape file and generates its outputs.",
	Long:  "https://github.com/charmbracelet/vhs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "help for vhs")
	rootCmd.Flags().StringSliceP("output", "o", nil, "file name(s) of video output")
	rootCmd.Flags().BoolP("publish", "p", false, "publish your GIF to vhs.charm.sh and get a shareable URL")
	rootCmd.Flags().BoolP("version", "v", false, "version for vhs")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet do not log messages. If publish flag is provided, it will log shareable URL")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".tape"),
	)
}
