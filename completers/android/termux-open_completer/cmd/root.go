package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-open",
	Short: "Open a file or URL in an external app",
	Long:  "https://github.com/termux/termux-tools/blob/master/scripts/termux-open.in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("chooser", false, "always show an app chooser dialog")
	rootCmd.Flags().String("content-type", "", "specify the MIME content `type` to use")
	rootCmd.Flags().BoolP("help", "h", false, "show usage and exit")
	rootCmd.Flags().Bool("send", false, "share the file for sending")
	rootCmd.Flags().Bool("view", false, "share the file for viewing (default)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"content-type": carapace.ActionValues(
			"application/json",
			"application/octet-stream",
			"application/pdf",
			"image/jpeg",
			"image/png",
			"text/html",
			"text/plain",
		),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
