package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "typst",
	Short: "A new markup-based typesetting system that is powerful and easy to learn.",
	Long:  "https://typst.app/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Print version")
	rootCmd.PersistentFlags().String("cert", "", "Path to a custom CA certificate to use when making network requests [env: TYPST_CERT=]")
	rootCmd.Flags().String("color", "auto", "Whether to use color")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cert":  carapace.ActionFiles(),
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
	})
}
