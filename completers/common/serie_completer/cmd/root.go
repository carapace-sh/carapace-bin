package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serie",
	Short: "A rich git commit graph in your terminal, like magic 📚",
	Long:  "https://github.com/lusingander/serie",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("graph-width", "g", "", "Commit graph image cell width")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringP("order", "o", "", "Commit ordering algorithm")
	rootCmd.Flags().Bool("preload", false, "Preload all graph images")
	rootCmd.Flags().StringP("protocol", "p", "", "Image protocol to render graph")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"graph-width": carapace.ActionValues("double", "single"),
		"order":       carapace.ActionValues("chrono", "topo"),
		"protocol":    carapace.ActionValues("auto", "iterm", "kitty").StyleF(style.ForKeyword),
	})
}
