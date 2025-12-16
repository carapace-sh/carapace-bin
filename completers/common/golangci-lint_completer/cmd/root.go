package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golangci-lint",
	Short: "golangci-lint is a smart linters runner.",
	Long:  "https://golangci-lint.run/",
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

	rootCmd.PersistentFlags().String("color", "", "Use color when printing; can be 'always', 'auto', or 'never'")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for a command")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.Flags().Bool("version", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword),
	})
}
